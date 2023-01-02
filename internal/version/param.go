package version

import (
	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateVersionParams(credentialID int64, versionDto *coreDto.VersionDto) db.CreateVersionParams {
	return db.CreateVersionParams{
		CredentialID: credentialID,
		Version:      *versionDto.Version,
		Url:          *versionDto.Url,
	}
}
