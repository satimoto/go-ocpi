package element

import "github.com/satimoto/go-datastore/db"

func NewCreateElementParams(dto *ElementDto) db.CreateElementParams {
	return db.CreateElementParams{}
}