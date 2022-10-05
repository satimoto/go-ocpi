package connector

import (
	"fmt"
	"strings"

	"github.com/satimoto/go-datastore/pkg/db"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/util"
)

func GetConnectorIdentifier(evse db.Evse, connectorDto *dto.ConnectorDto) *string {
	if evse.Identifier.Valid && connectorDto.Id != nil {
		if strings.Contains(evse.Identifier.String, "*") {
			identifier := fmt.Sprintf("%s*%s", util.TrimFromNthSeparator(evse.Identifier.String, 4, "*"), *connectorDto.Id)

			return &identifier
		}

		return &evse.Identifier.String
	}

	return nil
}
