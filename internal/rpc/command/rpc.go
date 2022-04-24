package command

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc/commandrpc"
)

func (r *RpcCommandResolver) ReserveNow(ctx context.Context, input *commandrpc.ReserveNowRequest) (*commandrpc.ReserveNowResponse, error) {
	if input != nil {
		expiryDate := util.ParseTime(input.ExpiryDate, nil)
		token, err := r.TokenResolver.Repository.GetTokenByUserID(ctx, db.GetTokenByUserIDParams{
			UserID: input.UserId,
			Type:   db.TokenTypeOTHER,
		})

		if err != nil {
			log.Printf("Error ReserveNow GetTokenByUserId: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Token not found")
		}

		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			log.Printf("Error ReserveNow GetLocationByUid: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			log.Printf("Error ReserveNow GetCredential: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Credential not found")
		}

		command, err := r.CommandResolver.ReserveNow(ctx, credential, token, location, &input.EvseUid, *expiryDate)

		if err != nil {
			log.Printf("Error ReserveNow ReserveNow: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Error requesting reservation")
		}

		reserveNowResponse := commandrpc.NewCommandReservationResponse(*command)

		return reserveNowResponse, nil
	}

	return nil, errors.New("Missing ReserveNowRequest")
}

func (r *RpcCommandResolver) StartSession(ctx context.Context, input *commandrpc.StartSessionRequest) (*commandrpc.StartSessionResponse, error) {
	if input != nil {
		token, err := r.TokenResolver.Repository.GetTokenByUserID(ctx, db.GetTokenByUserIDParams{
			UserID: input.UserId,
			Type:   db.TokenTypeOTHER,
		})

		if err != nil {
			log.Printf("Error StartSession GetTokenByUserId: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Token not found")
		}

		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			log.Printf("Error StartSession GetLocationByUid: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			log.Printf("Error StartSession GetCredential: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Credential not found")
		}

		command, err := r.CommandResolver.StartSession(ctx, credential, token, location, &input.EvseUid)

		if err != nil {
			log.Printf("Error StartSession StartSession: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Error starting session")
		}

		startResponse := commandrpc.NewCommandStartResponse(*command)

		return startResponse, nil
	}

	return nil, errors.New("Missing StartSessionRequest")
}

func (r *RpcCommandResolver) StopSession(ctx context.Context, input *commandrpc.StopSessionRequest) (*commandrpc.StopSessionResponse, error) {
	if input != nil {
		session, err := r.SessionResolver.Repository.GetSessionByUid(ctx, input.SessionUid)

		if err != nil {
			log.Printf("Error StopSession GetSessionByUid: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Token not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, session.CredentialID)

		if err != nil {
			log.Printf("Error StopSession GetCredential: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Credential not found")
		}

		command, err := r.CommandResolver.StopSession(ctx, credential, input.SessionUid)

		if err != nil {
			log.Printf("Error StopSession StopSession: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Error stopping session")
		}

		stopResponse := commandrpc.NewCommandStopResponse(*command)

		return stopResponse, nil
	}

	return nil, errors.New("Missing StopSessionRequest")
}

func (r *RpcCommandResolver) UnlockConnector(ctx context.Context, input *commandrpc.UnlockConnectorRequest) (*commandrpc.UnlockConnectorResponse, error) {
	if input != nil {
		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			log.Printf("Error UnlockConnector GetLocationByUid: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			log.Printf("Error UnlockConnector GetCredential: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Credential not found")
		}

		command, err := r.CommandResolver.UnlockConnector(ctx, credential, location, input.EvseUid, input.ConnectorUid)

		if err != nil {
			log.Printf("Error UnlockConnector UnlockConnector: %v", err)
			log.Printf("%#v", input)
			return nil, errors.New("Error unlocking connector")
		}

		unlockResponse := commandrpc.NewCommandUnlockResponse(*command)

		return unlockResponse, nil
	}

	return nil, errors.New("Missing StopSessionRequest")
}
