package location

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	coreLocation "github.com/satimoto/go-ocpi/internal/location"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *LocationResolver) SyncByIdentifier(ctx context.Context, credential db.Credential, fullSync bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	log.Printf("Sync locations Url=%v LastUpdated=%v CountryCode=%v PartyID=%v",
		credential.Url, lastUpdated, util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""))
	limit, offset, retries := 500, 0, 0

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, coreLocation.IDENTIFIER, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI125", "Error retrieving version endpoint", err)
		log.Printf("OCPI125: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, coreLocation.IDENTIFIER)
		return
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI126", "Error parsing url", err)
		log.Printf("OCPI126: Url=%v", versionEndpoint.Url)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, countryCode, partyID)
	query := requestUrl.Query()

	if fullSync {
		if countryCode != nil && partyID != nil {
			log.Printf("Setting all locations to removed")
			updateLocationsRemovedByPartyAndCountryCodeParams := db.UpdateLocationsRemovedByPartyAndCountryCodeParams{
				CountryCode: util.SqlNullString(countryCode),
				PartyID:     util.SqlNullString(partyID),
				IsRemoved:   true,
			}

			err := r.Repository.UpdateLocationsRemovedByPartyAndCountryCode(ctx, updateLocationsRemovedByPartyAndCountryCodeParams)

			if err != nil {
				metrics.RecordError("OCPI331", "Error updating locations", err)
				log.Printf("OCPI331: Param=%#v", updateLocationsRemovedByPartyAndCountryCodeParams)	
			}
		}
	} else {
		if lastUpdated != nil {
			query.Set("date_from", lastUpdated.UTC().Format(time.RFC3339))
		} else if location, err := r.GetLastLocationByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
			query.Set("date_from", location.LastUpdated.Format(time.RFC3339))
		}
	}

	for {
		if retries >= 5 {
			break
		}
		
		query.Set("limit", fmt.Sprintf("%d", limit))
		query.Set("offset", fmt.Sprintf("%d", offset))
		requestUrl.RawQuery = query.Encode()

		dto, limit := r.sendRequest(requestUrl.String(), header, limit)

		if dto == nil {
			retries++
			continue
		}

		retries = 0

		if dto.StatusCode == transportation.STATUS_CODE_OK {
			count := len(dto.Data)

			log.Printf("Page limit=%v offset=%v count=%v", limit, offset, count)
			r.ReplaceLocations(ctx, credential, dto.Data)
			offset += limit

			if limit == 0 || count == 0 || count < limit {
				break
			}
		}
	}
}

func (r *LocationResolver) sendRequest(url string, header transportation.OcpiRequestHeader, limit int) (*dto.OcpiLocationsDto, int) {
	response, err := r.OcpiService.Do(http.MethodGet, url, header, nil)

	if err != nil {
		metrics.RecordError("OCPI127", "Error making request", err)
		log.Printf("OCPI127: Method=%v, Url=%v, Header=%#v", http.MethodGet, url, header)
		
		return nil, limit
	}

	dto, err := r.UnmarshalPullDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		metrics.RecordError("OCPI128", "Error unmarshaling response", err)
		util.LogHttpResponse("OCPI128", url, response, true)
		
		return nil, limit
	}

	if dto.StatusCode != transportation.STATUS_CODE_OK {
		metrics.RecordError("OCPI129", "Error response failure", err)
		util.LogHttpRequest("OCPI129", url, response.Request, true)
		util.LogHttpResponse("OCPI129", url, response, true)
		log.Printf("OCPI129: StatusCode=%v, StatusMessage=%v", dto.StatusCode, dto.StatusMessage)
		
		return nil, limit
	}

	return dto, transportation.GetXLimitHeader(response, limit)
}
