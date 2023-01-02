package credential

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/util"
	ocpiVersion "github.com/satimoto/go-ocpi/internal/version"
)

func (r *CredentialResolver) PushCredential(ctx context.Context, httpMethod string, credential db.Credential) (*db.Credential, error) {
	identifier := "credentials"
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, identifier, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI264", "Error retrieving verion endpoint", err)
		log.Printf("OCPI264: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, identifier)
		return nil, transportation.OcpiRegistrationError(nil)
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	dto := r.CreateCredentialDto(ctx, credential)
	dtoBytes, err := json.Marshal(dto)

	if err != nil {
		metrics.RecordError("OCPI265", "Error marshalling dto", err)
		log.Printf("OCPI265: Dto=%#v", dto)
		return nil, transportation.OcpiRegistrationError(nil)
	}

	response, err := r.OcpiService.Do(httpMethod, versionEndpoint.Url, header, bytes.NewBuffer(dtoBytes))

	if err != nil {
		metrics.RecordError("OCPI266", "Error making request", err)
		log.Printf("OCPI266: Method=%v, Url=%v, Header=%#v", httpMethod, versionEndpoint.Url, header)
		return nil, transportation.OcpiRegistrationError(nil)
	}

	pullDto, err := r.UnmarshalPullDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		metrics.RecordError("OCPI267", "Error unmarshaling response", err)
		dbUtil.LogHttpResponse("OCPI267", versionEndpoint.Url, response, true)
		return nil, transportation.OcpiRegistrationError(nil)
	}

	if pullDto.StatusCode != transportation.STATUS_CODE_OK {
		metrics.RecordError("OCPI268", "Error response failure", err)
		dbUtil.LogHttpRequest("OCPI268", versionEndpoint.Url, response.Request, true)
		dbUtil.LogHttpResponse("OCPI268", versionEndpoint.Url, response, true)
		log.Printf("OCPI268: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
		return nil, transportation.OcpiRegistrationError(nil)
	}

	credentialDto := pullDto.Data
	updateCredentialParams := param.NewUpdateCredentialParams(credential)
	updateCredentialParams.LastUpdated = util.NewTimeUTC()

	if credentialDto.Token != nil {
		updateCredentialParams.ClientToken = dbUtil.SqlNullString(credentialDto.Token)
	}

	if credentialDto.Url != nil {
		updateCredentialParams.Url = *credentialDto.Url
	}

	if credentialDto.CountryCode != nil {
		updateCredentialParams.CountryCode = *credentialDto.CountryCode
	}

	if credentialDto.PartyID != nil {
		updateCredentialParams.PartyID = *credentialDto.PartyID
	}

	cred, err := r.Repository.UpdateCredential(ctx, updateCredentialParams)

	if err != nil {
		metrics.RecordError("OCPI269", "Error updating credential", err)
		log.Printf("OCPI269: Params=%#v", updateCredentialParams)
		return nil, transportation.OcpiRegistrationError(nil)
	}

	return &cred, nil
}

func (r *CredentialResolver) ReplaceCredential(ctx context.Context, credential db.Credential, credentialDto *dto.CredentialDto) (*db.Credential, error) {
	if credentialDto != nil {
		token := credential.ClientToken.String
		url := credential.Url
		partyID := credential.PartyID
		countryCode := credential.CountryCode

		if credentialDto.Token != nil {
			token = *credentialDto.Token
		}

		if credentialDto.Url != nil {
			url = *credentialDto.Url
		}

		if credentialDto.CountryCode != nil {
			countryCode = *credentialDto.CountryCode
		}

		if credentialDto.PartyID != nil {
			partyID = *credentialDto.PartyID
		}

		return r.RegisterCredential(ctx, credential, token, url, countryCode, partyID)
	}

	return nil, transportation.OcpiRegistrationError(nil)
}

func (r *CredentialResolver) RegisterCredential(ctx context.Context, credential db.Credential, token, url, countryCode, partyID string) (*db.Credential, error) {
	if len(token) > 0 {
		header := transportation.NewOcpiRequestHeader(&token, nil, nil)

		r.VersionResolver.PullVersions(ctx, url, header, credential.ID)
		version := r.VersionResolver.GetPreferredVersion(ctx, credential.ID, ocpiVersion.VERSION_2_1_1)

		if version != nil {
			r.VersionDetailResolver.PullVersionEndpoints(ctx, version.Url, header, version.ID)

			updateCredentialParams := param.NewUpdateCredentialParams(credential)
			updateCredentialParams.ClientToken = dbUtil.SqlNullString(token)
			updateCredentialParams.ServerToken = dbUtil.SqlNullString(uuid.NewString())
			updateCredentialParams.Url = url
			updateCredentialParams.CountryCode = countryCode
			updateCredentialParams.PartyID = partyID
			updateCredentialParams.VersionID = dbUtil.SqlNullInt64(version.ID)
			updateCredentialParams.LastUpdated = util.NewTimeUTC()

			cred, err := r.Repository.UpdateCredential(ctx, updateCredentialParams)

			if err != nil {
				metrics.RecordError("OCPI263", "Error updating credential", err)
				log.Printf("OCPI263: Params=%#v", updateCredentialParams)
				return nil, transportation.OcpiRegistrationError(nil)
			}

			go r.SyncService.SynchronizeCredential(cred, true, true, nil, nil, nil)

			return &cred, nil
		} else {
			return nil, transportation.OcpiUnsupportedVersion(nil)
		}
	}

	return nil, transportation.OcpiRegistrationError(nil)
}
