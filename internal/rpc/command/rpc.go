package command

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/ocpirpc"
	ocpiCommand "github.com/satimoto/go-ocpi/pkg/ocpi/command"
)

func (r *RpcCommandResolver) ReserveNow(reqCtx context.Context, input *ocpirpc.ReserveNowRequest) (*ocpirpc.ReserveNowResponse, error) {
	if input != nil {
		ctx := context.Background()
		expiryDate := util.ParseTime(input.ExpiryDate, nil)
		getTokenByUserIDParams := db.GetTokenByUserIDParams{
			UserID: input.UserId,
			Type:   db.TokenTypeOTHER,
		}
		token, err := r.TokenResolver.Repository.GetTokenByUserID(ctx, getTokenByUserIDParams)

		if err != nil {
			metrics.RecordError("OCPI140", "Error retrieving token", err)
			log.Printf("OCPI140: Params=%#v", getTokenByUserIDParams)
			return nil, errors.New("token not found")
		}

		if !token.Valid || token.Allowed != db.TokenAllowedTypeALLOWED {
			metrics.RecordError("OCPI141", "Error invalid token", err)
			log.Printf("OCPI141: Token=%#v", token)
			return nil, errors.New("token not found")
		}

		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			metrics.RecordError("OCPI142", "Error retrieving location", err)
			log.Printf("OCPI142: LocationUid=%v", input.LocationUid)
			return nil, errors.New("location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			metrics.RecordError("OCPI143", "Error retrieving credential", err)
			log.Printf("OCPI143: CredentialID=%v", location.CredentialID)
			return nil, errors.New("credential not found")
		}

		if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
			metrics.RecordError("OCPI144", "Error invalid credential", err)
			log.Printf("OCPI144: CredentialID=%v, Token=%v", credential.ID, credential.ClientToken)
			return nil, errors.New("error requesting reservation")
		}

		command, err := r.CommandResolver.ReserveNow(ctx, credential, token, location, &input.EvseUid, *expiryDate)

		if err != nil {
			metrics.RecordError("OCPI145", "Error requesting reservation", err)
			log.Printf("OCPI145: Input=%#v", input)
			return nil, err
		}

		reserveNowResponse := ocpiCommand.NewCommandReservationResponse(*command)

		return reserveNowResponse, nil
	}

	return nil, errors.New("missing request")
}

func (r *RpcCommandResolver) StartSession(reqCtx context.Context, input *ocpirpc.StartSessionRequest) (*ocpirpc.StartSessionResponse, error) {
	if input != nil {
		ctx := context.Background()
		getTokenByUserIDParams := db.GetTokenByUserIDParams{
			UserID: input.UserId,
			Type:   db.TokenTypeOTHER,
		}
		token, err := r.TokenResolver.Repository.GetTokenByUserID(ctx, getTokenByUserIDParams)

		if err != nil {
			metrics.RecordError("OCPI146", "Error retrieving token", err)
			log.Printf("OCPI146: Params=%#v", getTokenByUserIDParams)
			return nil, errors.New("token not found")
		}

		if !token.Valid || token.Allowed != db.TokenAllowedTypeALLOWED {
			metrics.RecordError("OCPI147", "Error invalid token", err)
			log.Printf("OCPI147: Token=%#v", token)
			return nil, errors.New("token not found")
		}

		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			metrics.RecordError("OCPI148", "Error retrieving location", err)
			log.Printf("OCPI148: LocationUid=%v", input.LocationUid)
			return nil, errors.New("location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			metrics.RecordError("OCPI149", "Error retrieving credential", err)
			log.Printf("OCPI149: CredentialID=%v", location.CredentialID)
			return nil, errors.New("credential not found")
		}

		if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
			metrics.RecordError("OCPI150", "Error invalid credential", err)
			log.Printf("OCPI150: CredentialID=%v, Token=%v", credential.ID, credential.ClientToken)
			return nil, errors.New("invalid credential token")
		}

		locationReferencesDto := dto.NewLocationReferencesDto(location.Uid)

		if len(input.EvseUid) > 0 {
			locationReferencesDto.EvseUids = []*string{&input.EvseUid}
		}

		tokenAuthorization, err := r.TokenResolver.TokenAuthorizationResolver.CreateTokenAuthorization(ctx, credential, token, locationReferencesDto)

		if err != nil {
			metrics.RecordError("OCPI151", "Error creating token authorization", err)
			log.Printf("OCPI151: LocationReferencesDto=%#v", locationReferencesDto)
			return nil, errors.New("error starting session")
		}

		command, err := r.CommandResolver.StartSession(ctx, credential, *tokenAuthorization, &input.EvseUid)

		if err != nil {
			metrics.RecordError("OCPI152", "Error requesting start", err)
			log.Printf("OCPI152: Input=%#v", input)
			return nil, err
		}

		startResponse := ocpiCommand.NewCommandStartResponse(*command)

		return startResponse, nil
	}

	return nil, errors.New("missing request")
}

func (r *RpcCommandResolver) StopSession(reqCtx context.Context, input *ocpirpc.StopSessionRequest) (*ocpirpc.StopSessionResponse, error) {
	if input != nil {
		ctx := context.Background()
		defaultResponse := ocpirpc.StopSessionResponse{
			Status:          string(db.CommandResponseTypeACCEPTED),
			AuthorizationId: input.AuthorizationId,
		}

		if tokenAuthorization, err := r.TokenResolver.TokenAuthorizationResolver.Repository.GetTokenAuthorizationByAuthorizationID(ctx, input.AuthorizationId); err == nil {
			updateTokenAuthorizationParams := param.NewUpdateTokenAuthorizationParams(tokenAuthorization)
			updateTokenAuthorizationParams.Authorized = false

			_, err := r.TokenResolver.TokenAuthorizationResolver.Repository.UpdateTokenAuthorizationByAuthorizationID(ctx, updateTokenAuthorizationParams)

			if err != nil {
				metrics.RecordError("OCPI153", "Error updating token authorization", err)
				log.Printf("OCPI153: Params=%#v", updateTokenAuthorizationParams)
			}
		}

		if session, err := r.SessionResolver.Repository.GetSessionByAuthorizationID(ctx, input.AuthorizationId); err == nil {
			if session.Uid == session.AuthorizationID.String || session.Status == db.SessionStatusTypeACTIVE || session.Status == db.SessionStatusTypePENDING {
				// This was a manually created session or the status is active/pending
				updateSessionByUidParams := param.NewUpdateSessionByUidParams(session)
				updateSessionByUidParams.Status = db.SessionStatusTypeINVALID

				if session.Status == db.SessionStatusTypeACTIVE {
					updateSessionByUidParams.Status = db.SessionStatusTypeCOMPLETED
				}

				updatedSession, err := r.SessionResolver.Repository.UpdateSessionByUid(ctx, updateSessionByUidParams)

				if err != nil {
					metrics.RecordError("OCPI309", "Error updating session", err)
					log.Printf("OCPI309: Params=%#v", updateSessionByUidParams)
				} else {
					sessionUpdateParams := param.NewCreateSessionUpdateParams(updatedSession)
					
					if _, err := r.SessionResolver.Repository.CreateSessionUpdate(ctx, sessionUpdateParams); err != nil {
						metrics.RecordError("OCPI338", "Error creating session update", err)
						log.Printf("OCPI338: Params=%#v", sessionUpdateParams)

					}
				}

				if session.Uid == session.AuthorizationID.String {
					return &defaultResponse, nil
				}
			}

			credential, err := r.CredentialResolver.Repository.GetCredential(ctx, session.CredentialID)

			if err != nil {
				metrics.RecordError("OCPI154", "Error retrieving credential", err)
				log.Printf("OCPI154: CredentialID=%v", session.CredentialID)
				return nil, errors.New("credential not found")
			}

			if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
				metrics.RecordError("OCPI155", "Error invalid credential", err)
				log.Printf("OCPI155: CredentialID=%v, Token=%v", credential.ID, credential.ClientToken)
				return nil, errors.New("invalid credential token")
			}

			command, err := r.CommandResolver.StopSession(ctx, credential, session.Uid)

			if err != nil {
				metrics.RecordError("OCPI156", "Error requesting stop", err)
				log.Printf("OCPI156: Input=%#v", input)
				return nil, err
			}

			stopResponse := ocpiCommand.NewCommandStopResponse(*command)

			return stopResponse, nil
		}

		return &defaultResponse, nil
	}

	return nil, errors.New("missing request")
}

func (r *RpcCommandResolver) UnlockConnector(reqCtx context.Context, input *ocpirpc.UnlockConnectorRequest) (*ocpirpc.UnlockConnectorResponse, error) {
	if input != nil {
		ctx := context.Background()
		location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, input.LocationUid)

		if err != nil {
			metrics.RecordError("OCPI157", "Error retrieving session", err)
			log.Printf("OCPI157: LocationUid=%v", input.LocationUid)
			return nil, errors.New("location not found")
		}

		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, location.CredentialID)

		if err != nil {
			metrics.RecordError("OCPI158", "Error retrieving credential", err)
			log.Printf("OCPI158: CredentialID=%v", location.CredentialID)
			return nil, errors.New("credential not found")
		}

		if !credential.ClientToken.Valid || len(credential.ClientToken.String) == 0 {
			metrics.RecordError("OCPI159", "Error invalid credential", err)
			log.Printf("OCPI159: CredentialID=%v, Token=%v", credential.ID, credential.ClientToken)
			return nil, errors.New("error requesting reservation")
		}

		command, err := r.CommandResolver.UnlockConnector(ctx, credential, location, input.EvseUid, input.ConnectorUid)

		if err != nil {
			metrics.RecordError("OCPI160", "Error requesting unlock", err)
			log.Printf("OCPI160: Input=%#v", input)
			return nil, err
		}

		unlockResponse := ocpiCommand.NewCommandUnlockResponse(*command)

		return unlockResponse, nil
	}

	return nil, errors.New("missing request")
}
