package credential

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *CredentialResolver) RegisterCredential(ctx context.Context, credential db.Credential, token, url, countryCode, partyID string) (*db.Credential, error) {
	if len(token) > 0 {
		header := transportation.NewOcpiRequestHeader(&token, nil, nil)

		r.VersionResolver.PullVersions(ctx, url, header, credential.ID)
		version := r.VersionResolver.GetBestVersion(ctx, credential.ID)

		if version != nil {
			r.VersionDetailResolver.PullVersionEndpoints(ctx, version.Url, header, version.ID)

			updateCredentialParams := param.NewUpdateCredentialParams(credential)
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
	if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
		log.Printf("OCPI010: Error credential not registered")
		log.Printf("OCPI010: CredentialID=%v, ClientToken=%v", credential.ID, credential.ClientToken)
		return nil, errors.New("error credential not registered")
	}

	if credential.ServerToken.Valid {
		versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "credentials", credential.CountryCode, credential.PartyID)

		if err != nil {
			util.LogOnError("OCPI011", "Error retrieving version endpoint", err)
			log.Printf("OCPI011: CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
			return nil, errors.New("error retrieving version endpoint")
		}

		updateCredentialParams := param.NewUpdateCredentialParams(credential)
		updateCredentialParams.ServerToken = util.SqlNullString(nil)

		cred, err := r.Repository.UpdateCredential(ctx, updateCredentialParams)

		if err != nil {
			util.LogOnError("OCPI012", "Error updating credential", err)
			log.Printf("OCPI012: Params=%#v", updateCredentialParams)
			return nil, errors.New("error updating credential")
		}

		header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
		response, err := r.OcpiRequester.Do(http.MethodDelete, versionEndpoint.Url, header, nil)

		if err != nil {
			util.LogOnError("OCPI013", "Error sending delete request", err)
			log.Printf("OCPI013: Url=%v", versionEndpoint.Url)
			return nil, errors.New("error sending delete request")
		}

		defer response.Body.Close()
		responseDto, err := transportation.UnmarshalResponseDto(response.Body)

		if err != nil {
			util.LogOnError("OCPI014", "Error unmarshalling response", err)
			util.LogHttpResponse("OCPI014", versionEndpoint.Url, response, true)
			return nil, errors.New("error unmarshalling response")
		}

		if responseDto.StatusCode != transportation.STATUS_CODE_OK {
			log.Printf("OCPI015: Error in delete response")
			log.Printf("OCPI015: Response=%#v", responseDto)
			return nil, errors.New("error in delete response")
		}

		return &cred, nil
	}

	return &credential, nil
}
