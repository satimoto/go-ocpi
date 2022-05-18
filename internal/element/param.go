package element

import "github.com/satimoto/go-datastore/pkg/db"

func NewCreateElementParams(dto *ElementDto) db.CreateElementParams {
	return db.CreateElementParams{}
}
