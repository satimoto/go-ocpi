package tokenauthorization

import (
	"context"
	"errors"

	"github.com/satimoto/go-ocpi/internal/async"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func (r *RpcTokenAuthorizationResolver) UpdateTokenAuthorization(ctx context.Context, request *ocpirpc.UpdateTokenAuthorizationRequest) (*ocpirpc.UpdateTokenAuthorizationResponse, error) {
	if request != nil {
		asyncResult := async.AsyncResult{
			String: request.AuthorizationId,
			Bool: request.Authorize,
		}

		ok := r.AsyncService.Set(request.AuthorizationId, asyncResult)

		return &ocpirpc.UpdateTokenAuthorizationResponse{Ok: ok}, nil
	}

	return nil, errors.New("error updating tokens")
}
