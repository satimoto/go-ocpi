package command

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/util"
	ocpiCommand "github.com/satimoto/go-ocpi/pkg/ocpi/command"
)

func (r *CommandResolver) ReserveNow(ctx context.Context, credential db.Credential, token db.Token, location db.Location, evseUid *string, expiryDate time.Time) (*db.CommandReservation, error) {
	identifier := "commands"
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, identifier, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI042", "Error retrieving version endpoint", err)
		log.Printf("OCPI042: CountryCode=%v, PartyID=%v, Idenitifer=%v", credential.CountryCode, credential.PartyID, identifier)
		return nil, errors.New("error requesting reservation")
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI043", "Error paring url", err)
		log.Printf("OCPI043: Url=%v", versionEndpoint.Url)
		return nil, errors.New("error requesting reservation")
	}

	createCommandReservationParams := ocpiCommand.NewCreateCommandReservationParams(token, expiryDate, location, evseUid)
	command, err := r.Repository.CreateCommandReservation(ctx, createCommandReservationParams)

	if err != nil {
		metrics.RecordError("OCPI044", "Error creating command reservation", err)
		log.Printf("OCPI044: Params=%#v", createCommandReservationParams)
		return nil, errors.New("error requesting reservation")
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	commandDto := r.CreateCommandReservationDto(ctx, command)
	dtoBytes, err := json.Marshal(commandDto)

	if err != nil {
		metrics.RecordError("OCPI045", "Error marshalling dto", err)
		log.Printf("OCPI045: Dto=%#v", commandDto)
		return nil, errors.New("error requesting reservation")
	}

	util.AppendPath(requestUrl, "RESERVE_NOW")
	response, err := r.OcpiService.Do(http.MethodPost, requestUrl.String(), header, bytes.NewBuffer(dtoBytes))

	if err != nil {
		metrics.RecordError("OCPI046", "Error making request", err)
		log.Printf("OCPI046: Method=%v, Url=%v, Header=%#v", http.MethodPost, requestUrl.String(), header)
		return nil, errors.New("error requesting reservation")
	}

	pullDto, err := r.UnmarshalPullDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		metrics.RecordError("OCPI047", "Error unmarshalling response", err)
		dbUtil.LogHttpResponse("OCPI047", requestUrl.String(), response, true)
		return nil, errors.New("error requesting reservation")
	}

	if pullDto.StatusCode != transportation.STATUS_CODE_OK {
		metrics.RecordError("OCPI047", "Error response failure", err)
		dbUtil.LogHttpRequest("OCPI048", requestUrl.String(), response.Request, true)
		dbUtil.LogHttpResponse("OCPI048", requestUrl.String(), response, true)
		log.Printf("OCPI048: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)

		updateCommandReservationParams := param.NewUpdateCommandReservationParams(command)
		updateCommandReservationParams.Status = db.CommandResponseTypeREJECTED
		r.Repository.UpdateCommandReservation(ctx, updateCommandReservationParams)

		return nil, errors.New("error requesting reservation")
	}

	if pullDto.Data.Result != nil && *pullDto.Data.Result != db.CommandResponseTypeACCEPTED {
		updateCommandReservationParams := param.NewUpdateCommandReservationParams(command)
		updateCommandReservationParams.Status = *pullDto.Data.Result

		if command, err = r.Repository.UpdateCommandReservation(ctx, updateCommandReservationParams); err == nil {
			return &command, nil
		}
	}

	return &command, nil
}

func (r *CommandResolver) StartSession(ctx context.Context, credential db.Credential, tokenAuthorization db.TokenAuthorization, evseUid *string) (*db.CommandStart, error) {
	identifier := "commands"
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, identifier, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI049", "Error retrieving version endpoint", err)
		log.Printf("OCPI049: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, identifier)
		return nil, errors.New("error starting session")
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI050", "Error paring url", err)
		log.Printf("OCPI050: Url=%v", versionEndpoint.Url)
		return nil, errors.New("error starting session")
	}

	createCommandStartParams := ocpiCommand.NewCreateCommandStartParams(tokenAuthorization, evseUid)
	command, err := r.Repository.CreateCommandStart(ctx, createCommandStartParams)

	if err != nil {
		metrics.RecordError("OCPI051", "Error creating command reservation", err)
		log.Printf("OCPI051: Params=%#v", createCommandStartParams)
		return nil, errors.New("error starting session")
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	commandDto := r.CreateCommandStartDto(ctx, command)
	dtoBytes, err := json.Marshal(commandDto)

	if err != nil {
		metrics.RecordError("OCPI052", "Error marshalling dto", err)
		log.Printf("OCPI052: Dto=%#v", commandDto)
		return nil, errors.New("error starting session")
	}

	util.AppendPath(requestUrl, "START_SESSION")
	response, err := r.OcpiService.Do(http.MethodPost, requestUrl.String(), header, bytes.NewBuffer(dtoBytes))

	if err != nil {
		metrics.RecordError("OCPI053", "Error making request", err)
		log.Printf("OCPI053: Method=%v, Url=%v, Header=%#v", http.MethodPost, requestUrl.String(), header)
		return nil, errors.New("error starting session")
	}

	defer response.Body.Close()
	pullDto, err := r.UnmarshalPullDto(response.Body)

	if err != nil {
		metrics.RecordError("OCPI054", "Error unmarshalling response", err)
		dbUtil.LogHttpResponse("OCPI054", requestUrl.String(), response, true)
		return nil, errors.New("error starting reservation")
	}

	if pullDto.StatusCode != transportation.STATUS_CODE_OK {
		metrics.RecordError("OCPI055", "Error response failure", err)
		dbUtil.LogHttpRequest("OCPI055", requestUrl.String(), response.Request, true)
		dbUtil.LogHttpResponse("OCPI055", requestUrl.String(), response, true)
		log.Printf("OCPI055: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)

		updateCommandStartParams := param.NewUpdateCommandStartParams(command)
		updateCommandStartParams.Status = db.CommandResponseTypeREJECTED
		r.Repository.UpdateCommandStart(ctx, updateCommandStartParams)

		return nil, errors.New("error starting reservation")
	}

	if pullDto.Data.Result != nil && *pullDto.Data.Result != db.CommandResponseTypeACCEPTED {
		updateCommandStartParams := param.NewUpdateCommandStartParams(command)
		updateCommandStartParams.Status = *pullDto.Data.Result

		if command, err = r.Repository.UpdateCommandStart(ctx, updateCommandStartParams); err == nil {
			return &command, nil
		}
	}

	return &command, nil
}

func (r *CommandResolver) StopSession(ctx context.Context, credential db.Credential, sessionID string) (*db.CommandStop, error) {
	identifier := "commands"
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, identifier, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI056", "Error retrieving version endpoint", err)
		log.Printf("OCPI056: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, identifier)
		return nil, errors.New("error stopping session")
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI057", "Error paring url", err)
		log.Printf("OCPI057: Url=%v", versionEndpoint.Url)
		return nil, errors.New("error stopping session")
	}

	createCommandStopParams := ocpiCommand.NewCreateCommandStopParams(sessionID)
	command, err := r.Repository.CreateCommandStop(ctx, createCommandStopParams)

	if err != nil {
		metrics.RecordError("OCPI058", "Error creating command reservation", err)
		log.Printf("OCPI058: Params=%#v", createCommandStopParams)
		return nil, errors.New("error stopping session")
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	commandDto := r.CreateCommandStopDto(ctx, command)
	dtoBytes, err := json.Marshal(commandDto)

	if err != nil {
		metrics.RecordError("OCPI059", "Error marshalling dto", err)
		log.Printf("OCPI059: Dto=%#v", commandDto)
		return nil, errors.New("error stopping session")
	}

	util.AppendPath(requestUrl, "STOP_SESSION")
	response, err := r.OcpiService.Do(http.MethodPost, requestUrl.String(), header, bytes.NewBuffer(dtoBytes))

	if err != nil {
		metrics.RecordError("OCPI060", "Error making request", err)
		log.Printf("OCPI060: Method=%v, Url=%v, Header=%#v", http.MethodPost, requestUrl.String(), header)
		return nil, errors.New("error stopping session")
	}

	defer response.Body.Close()
	pullDto, err := r.UnmarshalPullDto(response.Body)

	if err != nil {
		metrics.RecordError("OCPI061", "Error unmarshalling response", err)
		dbUtil.LogHttpResponse("OCPI062", requestUrl.String(), response, true)
		return nil, errors.New("error stopping reservation")
	}

	if pullDto.StatusCode != transportation.STATUS_CODE_OK {
		metrics.RecordError("OCPI063", "Error response failure", err)
		dbUtil.LogHttpRequest("OCPI063", requestUrl.String(), response.Request, true)
		dbUtil.LogHttpResponse("OCPI063", requestUrl.String(), response, true)
		log.Printf("OCPI063: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)

		updateCommandStopParams := param.NewUpdateCommandStopParams(command)
		updateCommandStopParams.Status = db.CommandResponseTypeREJECTED
		r.Repository.UpdateCommandStop(ctx, updateCommandStopParams)

		return nil, errors.New("error stopping reservation")
	}

	if pullDto.Data.Result != nil && *pullDto.Data.Result != db.CommandResponseTypeACCEPTED {
		updateCommandStopParams := param.NewUpdateCommandStopParams(command)
		updateCommandStopParams.Status = *pullDto.Data.Result

		if command, err = r.Repository.UpdateCommandStop(ctx, updateCommandStopParams); err == nil {
			return &command, nil
		}
	}

	return &command, nil
}

func (r *CommandResolver) UnlockConnector(ctx context.Context, credential db.Credential, location db.Location, evseUid string, connectorID string) (*db.CommandUnlock, error) {
	identifier := "commands"
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, identifier, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI064", "Error retrieving version endpoint", err)
		log.Printf("OCPI064: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, identifier)
		return nil, errors.New("error unlocking connector")
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI065", "Error paring url", err)
		log.Printf("OCPI065: Url=%v", versionEndpoint.Url)
		return nil, errors.New("error unlocking connector")
	}

	createCommandUnlockParams := ocpiCommand.NewCreateCommandUnlockParams(location, evseUid, connectorID)
	command, err := r.Repository.CreateCommandUnlock(ctx, createCommandUnlockParams)

	if err != nil {
		metrics.RecordError("OCPI066", "Error creating command reservation", err)
		log.Printf("OCPI066: Params=%#v", createCommandUnlockParams)
		return nil, errors.New("error unlocking connector")
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	commandDto := r.CreateCommandUnlockDto(ctx, command)
	dtoBytes, err := json.Marshal(commandDto)

	if err != nil {
		metrics.RecordError("OCPI067", "Error marshalling dto", err)
		log.Printf("OCPI067: Dto=%#v", commandDto)
		return nil, errors.New("error unlocking connector")
	}

	util.AppendPath(requestUrl, "UNLOCK_CONNECTOR")
	response, err := r.OcpiService.Do(http.MethodPost, requestUrl.String(), header, bytes.NewBuffer(dtoBytes))

	if err != nil {
		metrics.RecordError("OCPI068", "Error making request", err)
		log.Printf("OCPI068: Method=%v, Url=%v, Header=%#v", http.MethodPost, requestUrl.String(), header)
		return nil, errors.New("error unlocking connector")
	}

	defer response.Body.Close()
	pullDto, err := r.UnmarshalPullDto(response.Body)

	if err != nil {
		metrics.RecordError("OCPI069", "Error unmarshalling response", err)
		dbUtil.LogHttpResponse("OCPI069", requestUrl.String(), response, true)
		return nil, errors.New("error unlocking reservation")
	}

	if pullDto.StatusCode != transportation.STATUS_CODE_OK {
		metrics.RecordError("OCPI070", "Error response failure", err)
		dbUtil.LogHttpRequest("OCPI070", requestUrl.String(), response.Request, true)
		dbUtil.LogHttpResponse("OCPI070", requestUrl.String(), response, true)
		log.Printf("OCPI070: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)

		updateCommandUnlockParams := param.NewUpdateCommandUnlockParams(command)
		updateCommandUnlockParams.Status = db.CommandResponseTypeREJECTED
		r.Repository.UpdateCommandUnlock(ctx, updateCommandUnlockParams)

		return nil, errors.New("error unlocking reservation")
	}

	if pullDto.Data.Result != nil && *pullDto.Data.Result != db.CommandResponseTypeACCEPTED {
		updateCommandUnlockParams := param.NewUpdateCommandUnlockParams(command)
		updateCommandUnlockParams.Status = *pullDto.Data.Result

		if command, err = r.Repository.UpdateCommandUnlock(ctx, updateCommandUnlockParams); err == nil {
			return &command, nil
		}
	}

	return &command, nil
}
