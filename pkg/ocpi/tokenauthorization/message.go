package tokenauthorization

import (
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewUpdateTokenAuthorizationRequest(authorizationID string, authorized bool) *ocpirpc.UpdateTokenAuthorizationRequest {
	return &ocpirpc.UpdateTokenAuthorizationRequest{
		AuthorizationId: authorizationID,
		Authorize:       authorized,
	}
}
