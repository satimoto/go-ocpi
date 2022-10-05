package displaytext

import (
	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateDisplayTextParams(displayTextDto *coreDto.DisplayTextDto) db.CreateDisplayTextParams {
	return db.CreateDisplayTextParams{
		Language: displayTextDto.Language,
		Text:     displayTextDto.Text,
	}
}
