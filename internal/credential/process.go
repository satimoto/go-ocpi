package credential

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *CredentialResolver) RegisterCredential(ctx context.Context, credential db.Credential, token, url, countryCode, partyID string) (*db.Credential, error) {
	if len(token) > 0 {
		header := transportation.NewOcpiRequestHeader(&token, nil, nil)

		r.VersionResolver.PullVersions(ctx, url, header, credential.ID)
		version := r.VersionResolver.GetBestVersion(ctx, credential.ID)

		if version != nil {
			r.VersionDetailResolver.PullVersionEndpoints(ctx, version.Url, header, version.ID)

			updateCredentialParams := NewUpdateCredentialParams(credential)
			updateCredentialParams.ClientToken = util.SqlNullString(token)
			updateCredentialParams.ServerToken = util.SqlNullString(uuid.NewString())
			updateCredentialParams.Url = url
			updateCredentialParams.CountryCode = countryCode
			updateCredentialParams.PartyID = partyID
			updateCredentialParams.LastUpdated = time.Now()

			cred, err := r.Repository.UpdateCredential(ctx, updateCredentialParams)

			if err != nil {
				util.LogOnError("OCPI009", "Error updating credential", err)
				log.Printf("OCPI009: Params=%#v", updateCredentialParams)
				return nil, transportation.OcpiRegistrationError(nil)
			}

			go r.SyncResolver.SynchronizeCredential(ctx, cred)

			return &cred, nil
		} else {
			return nil, transportation.OcpiUnsupportedVersion(nil)
		}
	}

	return nil, transportation.OcpiRegistrationError(nil)
}

func (r *CredentialResolver) UnregisterCredential(ctx context.Context, credential db.Credential) (*db.Credential, error) {
	if credential.ServerToken.Valid {
		versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "credentials", credential.CountryCode, credential.PartyID)

		if err != nil {
			util.LogOnError("OCPI010", "Error retreiving version endpoint", err)
			log.Printf("OCPI010: CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
			return nil, errors.New("Error unregistering credential")
		}

		updateCredentialParams := NewUpdateCredentialParams(credential)
		updateCredentialParams.ServerToken = util.SqlNullString(nil)

		cred, err := r.Repository.UpdateCredential(ctx, updateCredentialParams)

		if err != nil {
			util.LogOnError("OCPI011", "Error updating credential", err)
			log.Printf("OCPI011: Params=%#v", updateCredentialParams)
			return nil, errors.New("Error unregistering credential")
		}

		header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
		response, err := r.OcpiRequester.Do(http.MethodDelete, versionEndpoint.Url, header, nil)

		if err != nil {
			util.LogOnError("OCPI012", "Error sending delete request", err)
			log.Printf("OCPI012: Url=%v", versionEndpoint.Url)
			return nil, errors.New("Error unregistering credential")
		}

		defer response.Body.Close()
		responseDto, err := transportation.UnmarshalResponseDto(response.Body)

		if err != nil {
			util.LogOnError("OCPI013", "Error unmarshalling response", err)
			util.LogHttpResponse("OCPI013", versionEndpoint.Url, response, true)
			return nil, errors.New("Error unregistering credential")
		}

		if responseDto.StatusCode != transportation.STATUS_CODE_OK {
			log.Printf("OCPI014: Error in delete request response")
			log.Printf("OCPI014: Response=%#v", responseDto)
			return nil, errors.New("Error unregistering credential")
		}

		return &cred, nil
	}

	return &credential, nil
}
