package displaytext

import "github.com/satimoto/go-datastore/db"

func NewCreateDisplayTextParams(dto *DisplayTextDto) db.CreateDisplayTextParams {
	return db.CreateDisplayTextParams{
		Language: dto.Language,
		Text:     dto.Text,
	}
}