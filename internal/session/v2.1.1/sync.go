package session

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	coreSession "github.com/satimoto/go-ocpi/internal/session"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *SessionResolver) SyncByIdentifier(ctx context.Context, credential db.Credential, fullSync bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	log.Printf("Sync sessions Url=%v LastUpdated=%v CountryCode=%v PartyID=%v",
		credential.Url, lastUpdated, util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""))
	limit, offset, retries := 500, 0, 0

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, coreSession.IDENTIFIER, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI170", "Error retrieving version endpoint", err)
		log.Printf("OCPI170: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, coreSession.IDENTIFIER)
		return
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI171", "Error parsing url", err)
		log.Printf("OCPI171: Url=%v", versionEndpoint.Url)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, countryCode, partyID)
	query := requestUrl.Query()

	if lastUpdated == nil {
		if session, err := r.GetLastSessionByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
			lastUpdated = &session.LastUpdated
		} else if credential.IsHub {
			oneMonthAgo := time.Now().Add(time.Hour * 24 * -30)
			lastUpdated = &oneMonthAgo
		}
	}

	if lastUpdated != nil {
		oneMonthFromLastUpdated := lastUpdated.Add(time.Hour * 24 * 30)

		if credential.IsHub && oneMonthFromLastUpdated.Before(time.Now()) {
			query.Set("date_to", oneMonthFromLastUpdated.Format(time.RFC3339))
		}

		query.Set("date_from", lastUpdated.Format(time.RFC3339))
	}

	for {
		query.Set("limit", fmt.Sprintf("%d", limit))
		query.Set("offset", fmt.Sprintf("%d", offset))
		requestUrl.RawQuery = query.Encode()

		response, err := r.OcpiService.Do(http.MethodGet, requestUrl.String(), header, nil)

		if err != nil {
			metrics.RecordError("OCPI172", "Error making request", err)
			log.Printf("OCPI172: Method=%v, Url=%v, Header=%#v", http.MethodGet, requestUrl.String(), header)
			retries++

			if retries >= 5 {
				break
			}

			continue
		}

		dto, err := r.UnmarshalPullDto(response.Body)
		defer response.Body.Close()

		if err != nil {
			metrics.RecordError("OCPI173", "Error unmarshaling response", err)
			util.LogHttpResponse("OCPI173", requestUrl.String(), response, true)
			break
		}

		limit = transportation.GetXLimitHeader(response, limit)

		if dto.StatusCode != transportation.STATUS_CODE_OK {
			metrics.RecordError("OCPI174", "Error response failure", err)
			util.LogHttpRequest("OCPI174", requestUrl.String(), response.Request, true)
			util.LogHttpResponse("OCPI174", requestUrl.String(), response, true)
			log.Printf("OCPI174: StatusCode=%v, StatusMessage=%v", dto.StatusCode, dto.StatusMessage)
			break
		}

		retries = 0

		if dto.StatusCode == transportation.STATUS_CODE_OK {
			count := len(dto.Data)

			log.Printf("Page limit=%v offset=%v count=%v", limit, offset, count)
			r.ReplaceSessions(ctx, credential, dto.Data)
			offset += limit

			if limit == 0 || count == 0 || count < limit {
				break
			}
		}
	}
}
