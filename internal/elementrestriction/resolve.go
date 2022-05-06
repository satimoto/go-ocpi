package elementrestriction

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type ElementRestrictionRepository interface {
	CreateElementRestriction(ctx context.Context, arg db.CreateElementRestrictionParams) (db.ElementRestriction, error)
	DeleteElementRestrictions(ctx context.Context, tariffID int64) error
	GetElementRestriction(ctx context.Context, id int64) (db.ElementRestriction, error)
	ListElementRestrictionWeekdays(ctx context.Context, elementRestrictionID int64) ([]db.Weekday, error)
	ListWeekdays(ctx context.Context) ([]db.Weekday, error)
	SetElementRestrictionWeekday(ctx context.Context, arg db.SetElementRestrictionWeekdayParams) error
	UnsetElementRestrictionWeekdays(ctx context.Context, elementRestrictionID int64) error
	UpdateElementRestriction(ctx context.Context, arg db.UpdateElementRestrictionParams) (db.ElementRestriction, error)
}

type ElementRestrictionResolver struct {
	Repository ElementRestrictionRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ElementRestrictionResolver {
	repo := ElementRestrictionRepository(repositoryService)
	
	return &ElementRestrictionResolver{repo}
}
