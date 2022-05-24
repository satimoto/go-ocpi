package token

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/pkg/evid"
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

		if _, err := r.Repository.GetTokenByAuthID(ctx, evIdValue); err != nil {
			return evIdValue, nil
		}

		attempts++

		if attempts > 1000000 {
			break
		}
	}

	log.Print("OCPI194", "Error generating AuthID")
	log.Printf("OCPI194: CountryCode=%v, PartyID=%v", countryCode, partyId)
	return "", errors.New("error generating AuthID")
}

func (r *TokenResolver) PushToken(ctx context.Context, httpMethod string, uid string, dto *TokenDto) {
	credentials, err := r.Repository.ListCredentials(ctx)

	if err != nil {
		dbUtil.LogOnError("OCPI195", "Error listing credentials", err)
		return
	}

	partyID := os.Getenv("API_PARTY_ID")
	countryCode := os.Getenv("API_COUNTRY_CODE")

	for _, credential := range credentials {
		if credential.ClientToken.Valid {
			versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "tokens", credential.CountryCode, credential.PartyID)

			if err != nil {
				dbUtil.LogOnError("OCPI196", "Error retrieving verion endpoint", err)
				log.Printf("OCPI196: CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
				continue
			}

			requestUrl, err := url.Parse(versionEndpoint.Url)

			if err != nil {
				dbUtil.LogOnError("OCPI197", "Error parsing url", err)
				log.Printf("OCPI197: Url=%v", versionEndpoint.Url)
				continue
			}

			header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
			dtoBytes, err := json.Marshal(dto)

			if err != nil {
				dbUtil.LogOnError("OCPI198", "Error marshalling dto", err)
				log.Printf("OCPI198: Dto=%#v", dto)
				continue
			}

			util.AppendPath(requestUrl, fmt.Sprintf("%s/%s/%s", countryCode, partyID, uid))
			response, err := r.OcpiRequester.Do(httpMethod, requestUrl.String(), header, bytes.NewReader(dtoBytes))

			if err != nil {
				dbUtil.LogOnError("OCPI199", "Error making request", err)
				log.Printf("OCPI199: Method=%v, Url=%v, Header=%#v", httpMethod, requestUrl.String(), header)
				continue
			}

			pullDto, err := transportation.UnmarshalResponseDto(response.Body)
			defer response.Body.Close()

			if err != nil {
				dbUtil.LogOnError("OCPI200", "Error unmarshalling response", err)
				dbUtil.LogHttpResponse("OCPI200", requestUrl.String(), response, true)
				break
			}

			if pullDto.StatusCode != transportation.STATUS_CODE_OK {
				dbUtil.LogHttpResponse("OCPI201", requestUrl.String(), response, true)
				log.Printf("OCPI201: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
			}
		}
	}
}

func (r *TokenResolver) ReplaceToken(ctx context.Context, userId int64, tokenAllowed db.TokenAllowedType, uid string, dto *TokenDto) *db.Token {
	if dto != nil {
		token, err := r.Repository.GetTokenByUid(ctx, uid)

		if err == nil {
			tokenParams := param.NewUpdateTokenByUidParams(token)
			tokenParams.Allowed = tokenAllowed

			if dto.AuthID != nil {
				tokenParams.AuthID = *dto.AuthID
			}

			if dto.Issuer != nil {
				tokenParams.Issuer = *dto.Issuer
			}

			if dto.Language != nil {
				tokenParams.Language = dbUtil.SqlNullString(dto.Language)
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
				tokenParams.VisualNumber = dbUtil.SqlNullString(dto.VisualNumber)
			}

			if dto.Whitelist != nil {
				tokenParams.Whitelist = *dto.Whitelist
			}

			updatedToken, err := r.Repository.UpdateTokenByUid(ctx, tokenParams)

			if err != nil {
				dbUtil.LogOnError("OCPI202", "Error updating token", err)
				log.Printf("OCPI202: Params=%#v", tokenParams)
				return nil
			}

			token = updatedToken
			r.PushToken(ctx, http.MethodPatch, token.Uid, dto)
		} else {
			tokenParams := NewCreateTokenParams(dto)
			tokenParams.Allowed = tokenAllowed
			tokenParams.UserID = userId

			token, err = r.Repository.CreateToken(ctx, tokenParams)

			if err != nil {
				dbUtil.LogOnError("OCPI203", "Error creating token", err)
				log.Printf("OCPI203: Params=%#v", tokenParams)
				return nil
			}

			r.PushToken(ctx, http.MethodPut, token.Uid, dto)
		}

		return &token
	}

	return nil
}
