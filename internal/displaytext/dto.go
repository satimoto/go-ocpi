package displaytext

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *DisplayTextResolver) CreateDisplayTextDto(ctx context.Context, displayText db.DisplayText) *coreDto.DisplayTextDto {
	return coreDto.NewDisplayTextDto(displayText)
}

func (r *DisplayTextResolver) CreateDisplayTextListDto(ctx context.Context, displayTexts []db.DisplayText) []*coreDto.DisplayTextDto {
	list := []*coreDto.DisplayTextDto{}

	for _, displayText := range displayTexts {
		list = append(list, r.CreateDisplayTextDto(ctx, displayText))
	}
	
	return list
}
