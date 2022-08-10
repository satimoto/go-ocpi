package command

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	tokenauthorization "github.com/satimoto/go-ocpi/internal/tokenauthorization/v2.1.1"
	"github.com/satimoto/go-ocpi/ocpirpc"
	ocpiCommand "github.com/satimoto/go-ocpi/pkg/ocpi/command"
)

func (r *RpcCommandResolver) ReserveNow(ctx context.Context, input *ocpirpc.ReserveNowRequest) (*ocpirpc.ReserveNowResponse, error) {
	if input != nil {
		expiryDate := util.ParseTime(input.ExpiryDate, nil)
		getTokenByUserIDParams := db.GetTokenByUserIDParams{
			UserID: input.UserId,
			Type:   db.TokenTypeOTHER,
		}
		token, err := r.TokenResolver.Repository.GetTokenByUserID(ctx, getTokenByUserIDParams)

		if err != nil {
			util.LogOnError("OCPI140", "Error retrieving token", err)
			log.Printf("OCPI140: Params=%#v", getTokenByUserIDParams)
			return nil, errors.New("token not found")
		}

		if !token.Valid || token.Allowed != db.TokenAllowedTypeALLOWED {
			util.LogOnError("OCPI141", "Error invalid token", err)
			log.Printf("OCPI141: Token=%#v", token)
			return nil, errors.New("token not found")
		}

		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			util.LogOnError("OCPI142", "Error retrieving location", err)
			log.Printf("OCPI142: LocationUid=%v", input.LocationUid)
			return nil, errors.New("location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			util.LogOnError("OCPI143", "Error retrieving credential", err)
			log.Printf("OCPI143: CredentialID=%v", location.CredentialID)
			return nil, errors.New("credential not found")
		}

		if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
			util.LogOnError("OCPI144", "Error invalid credential", err)
			log.Printf("OCPI144: CredentialID=%v, Token=%v", credential.ID, credential.ClientToken)
			return nil, errors.New("error requesting reservation")
		}

		command, err := r.CommandResolver.ReserveNow(ctx, credential, token, location, &input.EvseUid, *expiryDate)

		if err != nil {
			util.LogOnError("OCPI145", "Error requesting reservation", err)
			log.Printf("OCPI145: Input=%#v", input)
			return nil, errors.New("error requesting reservation")
		}

		reserveNowResponse := ocpiCommand.NewCommandReservationResponse(*command)

		return reserveNowResponse, nil
	}

	return nil, errors.New("missing request")
}

func (r *RpcCommandResolver) StartSession(ctx context.Context, input *ocpirpc.StartSessionRequest) (*ocpirpc.StartSessionResponse, error) {
	if input != nil {
		getTokenByUserIDParams := db.GetTokenByUserIDParams{
			UserID: input.UserId,
			Type:   db.TokenTypeOTHER,
		}
		token, err := r.TokenResolver.Repository.GetTokenByUserID(ctx, getTokenByUserIDParams)

		if err != nil {
			util.LogOnError("OCPI146", "Error retrieving token", err)
			log.Printf("OCPI146: Params=%#v", getTokenByUserIDParams)
			return nil, errors.New("token not found")
		}

		if !token.Valid || token.Allowed != db.TokenAllowedTypeALLOWED {
			util.LogOnError("OCPI147", "Error invalid token", err)
			log.Printf("OCPI147: Token=%#v", token)
			return nil, errors.New("token not found")
		}

		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			util.LogOnError("OCPI148", "Error retrieving location", err)
			log.Printf("OCPI148: LocationUid=%v", input.LocationUid)
			return nil, errors.New("location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			util.LogOnError("OCPI149", "Error retrieving credential", err)
			log.Printf("OCPI149: CredentialID=%v", location.CredentialID)
			return nil, errors.New("credential not found")
		}

		if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
			util.LogOnError("OCPI150", "Error invalid credential", err)
			log.Printf("OCPI150: CredentialID=%v, Token=%v", credential.ID, credential.ClientToken)
			return nil, errors.New("invalid credential token")
		}

		locationReferencesDto := tokenauthorization.NewLocationReferencesDto(location.Uid)

		if len(input.EvseUid) > 0 {
			locationReferencesDto.EvseUids = []*string{&input.EvseUid}
		}

		tokenAuthorization, err := r.TokenResolver.TokenAuthorizationResolver.CreateTokenAuthorization(ctx, token, locationReferencesDto)

		if err != nil {
			util.LogOnError("OCPI151", "Error creating token authorization", err)
			log.Printf("OCPI151: LocationReferencesDto=%#v", locationReferencesDto)
			return nil, errors.New("error starting session")
		}

		command, err := r.CommandResolver.StartSession(ctx, credential, *tokenAuthorization, &input.EvseUid)

		if err != nil {
			util.LogOnError("OCPI152", "Error requesting start", err)
			log.Printf("OCPI152: Input=%#v", input)
			return nil, errors.New("error starting session")
		}

		startResponse := ocpiCommand.NewCommandStartResponse(*command)

		return startResponse, nil
	}

	return nil, errors.New("missing request")
}

func (r *RpcCommandResolver) StopSession(ctx context.Context, input *ocpirpc.StopSessionRequest) (*ocpirpc.StopSessionResponse, error) {
	if input != nil {
		session, err := r.SessionResolver.Repository.GetSessionByUid(ctx, input.SessionUid)

		if err != nil {
			util.LogOnError("OCPI153", "Error retrieving session", err)
			log.Printf("OCPI153: SessionUid=%v", input.SessionUid)
			return nil, errors.New("session not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, session.CredentialID)

		if err != nil {
			util.LogOnError("OCPI154", "Error retrieving credential", err)
			log.Printf("OCPI154: CredentialID=%v", session.CredentialID)
			return nil, errors.New("credential not found")
		}

		if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
			util.LogOnError("OCPI155", "Error invalid credential", err)
			log.Printf("OCPI155: CredentialID=%v, Token=%v", credential.ID, credential.ClientToken)
			return nil, errors.New("invalid credential token")
		}

		command, err := r.CommandResolver.StopSession(ctx, credential, input.SessionUid)

		if err != nil {
			util.LogOnError("OCPI156", "Error requesting stop", err)
			log.Printf("OCPI156: Input=%#v", input)
			return nil, errors.New("error stopping session")
		}

		stopResponse := ocpiCommand.NewCommandStopResponse(*command)

		return stopResponse, nil
	}

	return nil, errors.New("missing request")
}

func (r *RpcCommandResolver) UnlockConnector(ctx context.Context, input *ocpirpc.UnlockConnectorRequest) (*ocpirpc.UnlockConnectorResponse, error) {
	if input != nil {
		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			util.LogOnError("OCPI157", "Error retrieving session", err)
			log.Printf("OCPI57: LocationUid=%v", input.LocationUid)
			return nil, errors.New("location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			util.LogOnError("OCPI158", "Error retrieving credential", err)
			log.Printf("OCPI158: CredentialID=%v", location.CredentialID)
			return nil, errors.New("credential not found")
		}

		if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
			util.LogOnError("OCPI159", "Error invalid credential", err)
			log.Printf("OCPI159: CredentialID=%v, Token=%v", credential.ID, credential.ClientToken)
			return nil, errors.New("error requesting reservation")
		}

		command, err := r.CommandResolver.UnlockConnector(ctx, credential, location, input.EvseUid, input.ConnectorUid)

		if err != nil {
			util.LogOnError("OCPI160", "Error requesting unlock", err)
			log.Printf("OCPI160: Input=%#v", input)
			return nil, errors.New("error unlocking connector")
		}

		unlockResponse := ocpiCommand.NewCommandUnlockResponse(*command)

		return unlockResponse, nil
	}

	return nil, errors.New("missing request")
}
