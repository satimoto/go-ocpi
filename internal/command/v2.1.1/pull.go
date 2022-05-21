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
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/util"
	ocpiCommand "github.com/satimoto/go-ocpi-api/pkg/ocpi/command"
)

func (r *CommandResolver) ReserveNow(ctx context.Context, credential db.Credential, token db.Token, location db.Location, evseUid *string, expiryDate time.Time) (*db.CommandReservation, error) {
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "commands", credential.CountryCode, credential.PartyID)

	if err != nil {
		log.Printf("Error ReserveNow GetVersionEndpointByIdentity: %v", err)
		log.Printf("CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
		return nil, errors.New("Error requesting reservation")
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		log.Printf("Error ReserveNow Parse: %v", err)
		log.Printf("Url=%v", versionEndpoint.Url)
		return nil, errors.New("Error requesting reservation")
	}

	createCommandReservationParams := ocpiCommand.NewCreateCommandReservationParams(token, expiryDate, location, evseUid)
	command, err := r.Repository.CreateCommandReservation(ctx, createCommandReservationParams)

	if err != nil {
		log.Printf("Error ReserveNow CreateCommandReservation: %v", err)
		log.Printf("Params=%#v", createCommandReservationParams)
		return nil, errors.New("Error requesting reservation")
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	dto := NewCommandReservationDto(command)
	dtoBytes, err := json.Marshal(dto)

	if err != nil {
		log.Printf("Error ReserveNow Marshal: %v", err)
		log.Printf("Dto=%#v", dto)
		return nil, errors.New("Error requesting reservation")
	}

	util.AppendPath(requestUrl, "RESERVE_NOW")
	response, err := r.OcpiRequester.Do(http.MethodPost, requestUrl.String(), header, bytes.NewReader(dtoBytes))

	if err != nil {
		log.Printf("Error ReserveNow Do: %v", err)
		log.Printf("Url=%v", requestUrl.String())
		return nil, errors.New("Error requesting reservation")
	}

	defer response.Body.Close()
	pullDto, err := r.UnmarshalPullDto(response.Body)

	if err != nil || pullDto.StatusCode != transportation.STATUS_CODE_OK {
		log.Printf("Error ReserveNow UnmarshalPullDto: %v", err)
		log.Printf("StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
		return nil, errors.New("Error requesting reservation")
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

func (r *CommandResolver) StartSession(ctx context.Context, credential db.Credential, token db.Token, location db.Location, evseUid *string) (*db.CommandStart, error) {
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "commands", credential.CountryCode, credential.PartyID)

	if err != nil {
		log.Printf("Error StartSession GetVersionEndpointByIdentity: %v", err)
		log.Printf("CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
		return nil, errors.New("Error starting session")
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		log.Printf("Error StartSession Parse: %v", err)
		log.Printf("Url=%v", versionEndpoint.Url)
		return nil, errors.New("Error starting session")
	}

	createCommandStartParams := ocpiCommand.NewCreateCommandStartParams(token, location, evseUid)
	command, err := r.Repository.CreateCommandStart(ctx, createCommandStartParams)

	if err != nil {
		log.Printf("Error StartSession CreateCommandStart: %v", err)
		log.Printf("Params=%#v", createCommandStartParams)
		return nil, errors.New("Error starting session")
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	dto := NewCommandStartDto(command)
	dtoBytes, err := json.Marshal(dto)

	if err != nil {
		log.Printf("Error StartSession Marshal: %v", err)
		log.Printf("Dto=%#v", dto)
		return nil, errors.New("Error starting session")
	}

	util.AppendPath(requestUrl, "START_SESSION")
	response, err := r.OcpiRequester.Do(http.MethodPost, requestUrl.String(), header, bytes.NewReader(dtoBytes))

	if err != nil {
		log.Printf("Error StartSession Do: %v", err)
		log.Printf("Url=%v", requestUrl.String())
		return nil, errors.New("Error starting session")
	}

	defer response.Body.Close()
	pullDto, err := r.UnmarshalPullDto(response.Body)

	if err != nil || pullDto.StatusCode != transportation.STATUS_CODE_OK {
		log.Printf("Error StartSession UnmarshalPullDto: %v", err)
		log.Printf("StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
		return nil, errors.New("Error starting session")
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
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "commands", credential.CountryCode, credential.PartyID)

	if err != nil {
		log.Printf("Error StopSession GetVersionEndpointByIdentity: %v", err)
		log.Printf("CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
		return nil, errors.New("Error stopping session")
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		log.Printf("Error StopSession Parse: %v", err)
		log.Printf("Url=%v", versionEndpoint.Url)
		return nil, errors.New("Error stopping session")
	}

	createCommandStopParams := ocpiCommand.NewCreateCommandStopParams(sessionID)
	command, err := r.Repository.CreateCommandStop(ctx, createCommandStopParams)

	if err != nil {
		log.Printf("Error StopSession CreateCommandStop: %v", err)
		log.Printf("Params=%#v", createCommandStopParams)
		return nil, errors.New("Error stopping session")
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	dto := NewCommandStopDto(command)
	dtoBytes, err := json.Marshal(dto)

	if err != nil {
		log.Printf("Error StopSession Marshal: %v", err)
		log.Printf("Dto=%#v", dto)
		return nil, errors.New("Error stopping session")
	}

	util.AppendPath(requestUrl, "STOP_SESSION")
	response, err := r.OcpiRequester.Do(http.MethodPost, requestUrl.String(), header, bytes.NewReader(dtoBytes))

	if err != nil {
		log.Printf("Error StopSession Do: %v", err)
		log.Printf("Url=%v", requestUrl.String())
		return nil, errors.New("Error stopping session")
	}

	defer response.Body.Close()
	pullDto, err := r.UnmarshalPullDto(response.Body)

	if err != nil || pullDto.StatusCode != transportation.STATUS_CODE_OK {
		log.Printf("Error StopSession UnmarshalPullDto: %v", err)
		log.Printf("StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
		return nil, errors.New("Error stopping session")
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
	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "commands", credential.CountryCode, credential.PartyID)

	if err != nil {
		log.Printf("Error UnlockConnector GetVersionEndpointByIdentity: %v", err)
		log.Printf("CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
		return nil, errors.New("Error unlocking connector")
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		log.Printf("Error UnlockConnector Parse: %v", err)
		log.Printf("Url=%v", versionEndpoint.Url)
		return nil, errors.New("Error unlocking connector")
	}

	createCommandUnlockParams := ocpiCommand.NewCreateCommandUnlockParams(location, evseUid, connectorID)
	command, err := r.Repository.CreateCommandUnlock(ctx, createCommandUnlockParams)

	if err != nil {
		log.Printf("Error UnlockConnector CreateCommandUnlock: %v", err)
		log.Printf("Params=%#v", createCommandUnlockParams)
		return nil, errors.New("Error unlocking connector")
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)
	dto := NewCommandUnlockDto(command)
	dtoBytes, err := json.Marshal(dto)

	if err != nil {
		log.Printf("Error UnlockConnector Marshal: %v", err)
		log.Printf("Dto=%#v", dto)
		return nil, errors.New("Error unlocking connector")
	}

	util.AppendPath(requestUrl, "UNLOCK_CONNECTOR")
	response, err := r.OcpiRequester.Do(http.MethodPost, requestUrl.String(), header, bytes.NewReader(dtoBytes))

	if err != nil {
		log.Printf("Error UnlockConnector Do: %v", err)
		log.Printf("Url=%v", requestUrl.String())
		return nil, errors.New("Error unlocking connector")
	}

	defer response.Body.Close()
	pullDto, err := r.UnmarshalPullDto(response.Body)

	if err != nil || pullDto.StatusCode != transportation.STATUS_CODE_OK {
		log.Printf("Error UnlockConnector UnmarshalPullDto: %v", err)
		log.Printf("StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
		return nil, errors.New("Error unlocking connector")
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
