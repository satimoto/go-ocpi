package element

import (
	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateElementParams(elementDto *coreDto.ElementDto) db.CreateElementParams {
	return db.CreateElementParams{}
}
