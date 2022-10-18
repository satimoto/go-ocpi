package tokenauthorization

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/async"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func (r *RpcTokenAuthorizationResolver) UpdateTokenAuthorization(ctx context.Context, request *ocpirpc.UpdateTokenAuthorizationRequest) (*ocpirpc.UpdateTokenAuthorizationResponse, error) {
	if request != nil {
		tokenAuthorization, err := r.TokenAuthorizationRepository.GetTokenAuthorizationByAuthorizationID(ctx, request.AuthorizationId)

		if err != nil {
			util.LogOnError("OCPI289", "Error retrieving token authorization", err)
			log.Printf("OCPI289: Request=%#v", request)
			return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: false}, nil
		}

		token, err := r.TokenRepository.GetToken(ctx, tokenAuthorization.TokenID)

		if err != nil {
			util.LogOnError("OCPI290", "Error retrieving token", err)
			log.Printf("OCPI290: TokenID=%v", tokenAuthorization.TokenID)
			return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: false}, nil
		}

		if token.Type == db.TokenTypeRFID {
			// Update token authorization using async channel
			asyncResult := async.AsyncResult{
				String: request.AuthorizationId,
				Bool:   request.Authorize,
			}

			ok := r.AsyncService.Set(request.AuthorizationId, asyncResult)

			return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: ok}, nil
		} else if _, err := r.SessionRepository.GetSessionByAuthorizationID(ctx, tokenAuthorization.AuthorizationID); err != nil {
			// Only update token authorization if session is not yet created
			updateTokenAuthorizationParams := db.UpdateTokenAuthorizationByAuthorizationIDParams{
				AuthorizationID: request.AuthorizationId,
				Authorized:      request.Authorize,
			}

			_, err := r.TokenAuthorizationRepository.UpdateTokenAuthorizationByAuthorizationID(ctx, updateTokenAuthorizationParams)

			if err != nil {
				util.LogOnError("OCPI291", "Error updating token authorization", err)
				log.Printf("OCPI291: Params=%#v", updateTokenAuthorizationParams)
				return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: false}, nil
			}

			return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: true}, nil
		}
	}

	return nil, errors.New("error updating token authorization")
}
