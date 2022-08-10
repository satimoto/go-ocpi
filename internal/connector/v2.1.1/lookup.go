package connector

import (
	"fmt"
	"strings"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/util"
)

func GetConnectorIdentifier(evse db.Evse, dto *ConnectorDto) *string {
	if evse.Identifier.Valid && dto.Id != nil {
		if strings.Contains(evse.Identifier.String, "*") {
			identifier := fmt.Sprintf("%s*%s", util.TrimFromNthSeparator(evse.Identifier.String, 4, "*"), *dto.Id)

			return &identifier
		}

		return &evse.Identifier.String
	}

	return nil
}
