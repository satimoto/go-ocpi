package session

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func (r *SessionResolver) GetLastSessionByIdentity(ctx context.Context, countryCode string, partyID string) (db.Session, error) {
	params := db.GetSessionByIdentityOrderByLastUpdatedParams{
		CountryCode: util.SqlNullString(countryCode),
		PartyID:     util.SqlNullString(partyID),
	}

	return r.Repository.GetSessionByIdentityOrderByLastUpdated(ctx, params)
}
