package tokenauthorization

import (
	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateTokenAuthorizationParams(tokenID int64) db.CreateTokenAuthorizationParams {
	return db.CreateTokenAuthorizationParams{
		TokenID:         tokenID,
		AuthorizationID: uuid.NewString(),
	}
}

func NewUpdateTokenAuthorizationParams(authorizationID string, countryCode *string, partyID *string) db.UpdateTokenAuthorizationByAuthorizationIDParams {
	return db.UpdateTokenAuthorizationByAuthorizationIDParams{
		AuthorizationID: authorizationID,
		CountryCode:     util.SqlNullString(countryCode),
		PartyID:         util.SqlNullString(partyID),
	}
}
