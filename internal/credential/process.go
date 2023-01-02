package credential

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/util"
	ocpiVersion "github.com/satimoto/go-ocpi/internal/version"
)

func (r *CredentialResolver) RegisterCredential(ctx context.Context, credential db.Credential, token string) (*db.Credential, error) {
	if len(token) > 0 {
		header := transportation.NewOcpiRequestHeader(&token, nil, nil)

		r.VersionResolver.PullVersions(ctx, credential.Url, header, credential.ID)
		version := r.VersionResolver.GetBestVersion(ctx, credential.ID)

		if version != nil {
			r.VersionDetailResolver.PullVersionEndpoints(ctx, version.Url, header, version.ID)

			updateCredentialParams := param.NewUpdateCredentialParams(credential)
			updateCredentialParams.ClientToken = dbUtil.SqlNullString(token)
			updateCredentialParams.ServerToken = dbUtil.SqlNullString(uuid.NewString())
			updateCredentialParams.VersionID = dbUtil.SqlNullInt64(version.ID)
			updateCredentialParams.LastUpdated = util.NewTimeUTC()

			registeredCredential, err := r.Repository.UpdateCredential(ctx, updateCredentialParams)

			if err != nil {
				metrics.RecordError("OCPI009", "Error updating credential", err)
				log.Printf("OCPI009: Params=%#v", updateCredentialParams)
				return nil, transportation.OcpiRegistrationError(nil)
			}

			httpMethod := http.MethodPost

			if credential.ServerToken.Valid {
				httpMethod = http.MethodPut
			}

			if pushedCredential, err := r.PushCredential(ctx, httpMethod, *version, registeredCredential); err == nil {
				go r.SyncService.SynchronizeCredential(*pushedCredential, true, true, nil, nil, nil)

				return pushedCredential, nil
			}
		} else {
			return nil, transportation.OcpiUnsupportedVersion(nil)
		}
	}

	return nil, transportation.OcpiRegistrationError(nil)
}

func (r *CredentialResolver) PushCredential(ctx context.Context, httpMethod string, version db.Version, credential db.Credential) (*db.Credential, error) {
	if version.Version == ocpiVersion.VERSION_2_1_1 {
		return r.CredentialResolver_2_1_1.PushCredential(ctx, httpMethod, credential)
	}

	return nil, transportation.OcpiUnsupportedVersion(nil)
}

func (r *CredentialResolver) UnregisterCredential(ctx context.Context, credential db.Credential) (*db.Credential, error) {
	if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
		log.Printf("OCPI010: Error credential not registered")
		log.Printf("OCPI010: CredentialID=%v, ClientToken=%v", credential.ID, credential.ClientToken)
		return nil, errors.New("error credential not registered")
	}

	if credential.ServerToken.Valid {
		identifier := "credentials"
		versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, identifier, credential.CountryCode, credential.PartyID)

		if err != nil {
			metrics.RecordError("OCPI011", "Error retrieving version endpoint", err)
			log.Printf("OCPI011: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, identifier)
			return nil, errors.New("error retrieving version endpoint")
		}

		updateCredentialParams := param.NewUpdateCredentialParams(credential)
		updateCredentialParams.ServerToken = dbUtil.SqlNullString(nil)

		cred, err := r.Repository.UpdateCredential(ctx, updateCredentialParams)

		if err != nil {
			metrics.RecordError("OCPI012", "Error updating credential", err)
			log.Printf("OCPI012: Params=%#v", updateCredentialParams)
			return nil, errors.New("error updating credential")
		}

		header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
		response, err := r.OcpiService.Do(http.MethodDelete, versionEndpoint.Url, header, nil)

		if err != nil {
			metrics.RecordError("OCPI013", "Error sending delete request", err)
			log.Printf("OCPI013: Url=%v", versionEndpoint.Url)
			return nil, errors.New("error sending delete request")
		}

		defer response.Body.Close()
		responseDto, err := transportation.UnmarshalResponseDto(response.Body)

		if err != nil {
			metrics.RecordError("OCPI014", "Error unmarshaling response", err)
			dbUtil.LogHttpResponse("OCPI014", versionEndpoint.Url, response, true)
			return nil, errors.New("error unmarshaling response")
		}

		if responseDto.StatusCode != transportation.STATUS_CODE_OK {
			metrics.RecordError("OCPI015", "Error response failure", err)
			dbUtil.LogHttpRequest("OCPI015", versionEndpoint.Url, response.Request, true)
			dbUtil.LogHttpResponse("OCPI015", versionEndpoint.Url, response, true)
			log.Printf("OCPI015: StatusCode=%v, StatusMessage=%v", responseDto.StatusCode, responseDto.StatusMessage)
			return nil, errors.New("error in delete response")
		}

		return &cred, nil
	}

	return &credential, nil
}
