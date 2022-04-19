package token

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"os"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/internal/evid"
)

func (r *TokenResolver) GenerateAuthID(ctx context.Context) (string, error) {
	countryCode := os.Getenv("COUNTRY_CODE")
	partyId := os.Getenv("PARTY_ID")
	authBytes := make([]byte, 4)
	attempts := 0

	for {
		rand.Read(authBytes)
		evId := evid.NewEvid(fmt.Sprintf("%s*%s*C%x", countryCode, partyId, authBytes))
		evIdValue := evId.ValueWithSeparator("*")

		if _, err := r.Repository.GetTokenByAuthId(ctx, evIdValue); err != nil {
			return evIdValue, nil
		}

		attempts++

		if attempts > 1000000 {
			break
		}
	}

	return "", errors.New("Error generating AuthID")
}

func (r *TokenResolver) ReplaceToken(ctx context.Context, uid string, dto *TokenDto) *db.Token {
	if dto != nil {
		token, err := r.Repository.GetTokenByUid(ctx, uid)

		if err == nil {
			tokenParams := NewUpdateTokenByUidParams(token)

			if dto.AuthID != nil {
				tokenParams.AuthID = *dto.AuthID
			}

			if dto.Issuer != nil {
				tokenParams.Issuer = *dto.Issuer
			}

			if dto.Language != nil {
				tokenParams.Language = util.SqlNullString(dto.Language)
			}

			if dto.LastUpdated != nil {
				tokenParams.LastUpdated = *dto.LastUpdated
			}

			if dto.Type != nil {
				tokenParams.Type = *dto.Type
			}

			if dto.Valid != nil {
				tokenParams.Valid = *dto.Valid
			}

			if dto.VisualNumber != nil {
				tokenParams.VisualNumber = util.SqlNullString(dto.VisualNumber)
			}

			if dto.Whitelist != nil {
				tokenParams.Whitelist = *dto.Whitelist
			}

			token, err = r.Repository.UpdateTokenByUid(ctx, tokenParams)
		} else {
			tokenParams := NewCreateTokenParams(dto)

			token, err = r.Repository.CreateToken(ctx, tokenParams)
		}

		return &token
	}

	return nil
}
