package displaytext

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type DisplayTextRepository interface {
	CreateDisplayText(ctx context.Context, arg db.CreateDisplayTextParams) (db.DisplayText, error)
}

type DisplayTextResolver struct {
	Repository DisplayTextRepository
}

func NewResolver(repositoryService *db.RepositoryService) *DisplayTextResolver {
	repo := DisplayTextRepository(repositoryService)
	return &DisplayTextResolver{repo}
}
