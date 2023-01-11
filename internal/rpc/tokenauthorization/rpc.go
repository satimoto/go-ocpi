package tokenauthorization

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/async"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func (r *RpcTokenAuthorizationResolver) UpdateTokenAuthorization(reqCtx context.Context, request *ocpirpc.UpdateTokenAuthorizationRequest) (*ocpirpc.UpdateTokenAuthorizationResponse, error) {
	if request != nil {
		ctx := context.Background()
		tokenAuthorization, err := r.TokenAuthorizationRepository.GetTokenAuthorizationByAuthorizationID(ctx, request.AuthorizationId)

		if err != nil {
			metrics.RecordError("OCPI289", "Error retrieving token authorization", err)
			log.Printf("OCPI289: Request=%#v", request)
			return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: false}, nil
		}

		token, err := r.TokenRepository.GetToken(ctx, tokenAuthorization.TokenID)

		if err != nil {
			metrics.RecordError("OCPI290", "Error retrieving token", err)
			log.Printf("OCPI290: TokenID=%v", tokenAuthorization.TokenID)
			return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: false}, nil
		}

		log.Printf("Update token authorization: %v", request.AuthorizationId)

		if token.Type == db.TokenTypeRFID {
			// Update token authorization using async channel
			asyncResult := async.AsyncResult{
				String: request.AuthorizationId,
				Bool:   request.Authorize,
			}

			ok := r.AsyncService.Set(request.AuthorizationId, asyncResult)

			log.Printf("Async result authorize=%v ok=%v", request.Authorize, ok)

			return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: ok}, nil
		} else if _, err := r.SessionRepository.GetSessionByAuthorizationID(ctx, tokenAuthorization.AuthorizationID); err != nil {
			// Only update token authorization if session is not yet created
			updateTokenAuthorizationParams := db.UpdateTokenAuthorizationByAuthorizationIDParams{
				AuthorizationID: request.AuthorizationId,
				Authorized:      request.Authorize,
			}

			log.Printf("Token authorise=%v", request.Authorize)

			_, err := r.TokenAuthorizationRepository.UpdateTokenAuthorizationByAuthorizationID(ctx, updateTokenAuthorizationParams)

			if err != nil {
				metrics.RecordError("OCPI291", "Error updating token authorization", err)
				log.Printf("OCPI291: Params=%#v", updateTokenAuthorizationParams)
				return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: false}, nil
			}

			return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: true}, nil
		}
	}

	return nil, errors.New("error updating token authorization")
}
