package token

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
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
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/util"
	"github.com/satimoto/go-ocpi/pkg/evid"
)

func (r * TokenResolver) GenerateUid(ctx context.Context) (string, error) {
	uidBytes := make([]byte, 10)
	attempts := 0

	for {
		rand.Read(uidBytes)
		uid := hex.EncodeToString(uidBytes)

		if _, err := r.Repository.GetTokenByUid(ctx, uid); err != nil {
			return uid, nil
		}

		attempts++

		if attempts > 1000000 {
			break
		}
	}

	log.Print("OCPI339", "Error generating Uid")
	return "", errors.New("error generating Uid")

}

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

func (r *TokenResolver) PushToken(ctx context.Context, httpMethod string, uid string, tokenDto *dto.TokenDto) {
	credentials, err := r.Repository.ListCredentials(ctx)

	if err != nil {
		metrics.RecordError("OCPI195", "Error listing credentials", err)
		return
	}

	partyID := os.Getenv("PARTY_ID")
	countryCode := os.Getenv("COUNTRY_CODE")

	for _, credential := range credentials {
		if credential.ClientToken.Valid {
			versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "tokens", credential.CountryCode, credential.PartyID)

			if err != nil {
				metrics.RecordError("OCPI196", "Error retrieving verion endpoint", err)
				log.Printf("OCPI196: CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
				continue
			}

			requestUrl, err := url.Parse(versionEndpoint.Url)

			if err != nil {
				metrics.RecordError("OCPI197", "Error parsing url", err)
				log.Printf("OCPI197: Url=%v", versionEndpoint.Url)
				continue
			}

			header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
			dtoBytes, err := json.Marshal(tokenDto)

			if err != nil {
				metrics.RecordError("OCPI198", "Error marshalling dto", err)
				log.Printf("OCPI198: Dto=%#v", tokenDto)
				continue
			}

			util.AppendPath(requestUrl, fmt.Sprintf("%s/%s/%s", countryCode, partyID, uid))
			response, pullDto, err := r.sendRequest(requestUrl, httpMethod, header, dtoBytes)

			if response != nil && dtoBytes != nil {
				if pullDto.StatusCode != transportation.STATUS_CODE_OK {
					if pullDto.StatusCode == transportation.STATUS_CODE_REGISTRATION_ERROR && httpMethod == http.MethodPatch {
						response, pullDto, err = r.sendRequest(requestUrl, http.MethodPut, header, dtoBytes)
					}

					if response != nil && pullDto != nil && pullDto.StatusCode != transportation.STATUS_CODE_OK {
						metrics.RecordError("OCPI201", "Error response failure", err)
						dbUtil.LogHttpRequest("OCPI201", requestUrl.String(), response.Request, true)
						dbUtil.LogHttpResponse("OCPI201", requestUrl.String(), response, true)
						log.Printf("OCPI201: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
					}
				}

				response.Body.Close()
			}
		}
	}
}

func (r *TokenResolver) sendRequest(requestUrl *url.URL, httpMethod string, header transportation.OcpiRequestHeader, dtoBytes []byte) (*http.Response, *transportation.OcpiResponse, error) {
	response, err := r.OcpiService.Do(httpMethod, requestUrl.String(), header, bytes.NewBuffer(dtoBytes))

	if err != nil {
		metrics.RecordError("OCPI199", "Error making request", err)
		log.Printf("OCPI200: Method=%v, Url=%v, Header=%#v", httpMethod, requestUrl.String(), header)
		return response, nil, err
	}

	pullDto, err := transportation.UnmarshalResponseDto(response.Body)

	if err != nil {
		metrics.RecordError("OCPI200", "Error unmarshaling response", err)
		dbUtil.LogHttpResponse("OCPI200", requestUrl.String(), response, true)
		return response, nil, err
	}

	return response, pullDto, nil
}

func (r *TokenResolver) ReplaceToken(ctx context.Context, userId int64, tokenAllowed db.TokenAllowedType, uid string, tokenDto *dto.TokenDto) *db.Token {
	if tokenDto != nil {
		token, err := r.Repository.GetTokenByUid(ctx, uid)

		if err == nil {
			tokenParams := param.NewUpdateTokenByUidParams(token)
			tokenParams.Allowed = tokenAllowed

			if tokenDto.AuthID != nil {
				tokenParams.AuthID = *tokenDto.AuthID
			}

			if tokenDto.Issuer != nil {
				tokenParams.Issuer = *tokenDto.Issuer
			}

			if tokenDto.Language != nil {
				tokenParams.Language = dbUtil.SqlNullString(tokenDto.Language)
			}

			if tokenDto.LastUpdated != nil {
				tokenParams.LastUpdated = tokenDto.LastUpdated.Time()
			}

			if tokenDto.Type != nil {
				tokenParams.Type = *tokenDto.Type
			}

			if tokenDto.Valid != nil {
				tokenParams.Valid = *tokenDto.Valid
			}

			if tokenDto.VisualNumber != nil {
				tokenParams.VisualNumber = dbUtil.SqlNullString(tokenDto.VisualNumber)
			}

			if tokenDto.Whitelist != nil {
				tokenParams.Whitelist = *tokenDto.Whitelist
			}

			updatedToken, err := r.Repository.UpdateTokenByUid(ctx, tokenParams)

			if err != nil {
				metrics.RecordError("OCPI202", "Error updating token", err)
				log.Printf("OCPI202: Params=%#v", tokenParams)
				return nil
			}

			token = updatedToken
			r.PushToken(ctx, http.MethodPatch, token.Uid, tokenDto)
		} else {
			tokenParams := NewCreateTokenParams(tokenDto)
			tokenParams.Allowed = tokenAllowed
			tokenParams.UserID = userId

			token, err = r.Repository.CreateToken(ctx, tokenParams)

			if err != nil {
				metrics.RecordError("OCPI203", "Error creating token", err)
				log.Printf("OCPI203: Params=%#v", tokenParams)
				return nil
			}

			r.PushToken(ctx, http.MethodPut, token.Uid, tokenDto)
		}

		return &token
	}

	return nil
}
