package version

import (
	"github.com/satimoto/go-datastore/db"
)

func NewCreateVersionParams(credentialID int64, dto *VersionDto) db.CreateVersionParams {
	return db.CreateVersionParams{
		CredentialID: credentialID,
		Version:      *dto.Version,
		Url:          *dto.Url,
	}
}
