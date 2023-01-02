package connector_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	connectorMocks "github.com/satimoto/go-ocpi/internal/connector/v2.1.1/mocks"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestReplaceConnector(t *testing.T) {
	ctx := context.Background()

	t.Run("Create Connector", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		connectorResolver := connectorMocks.NewResolver(mockRepository)

		connectorTypeIEC62196T2 := db.ConnectorTypeIEC62196T2
		connectorFormatCABLE := db.ConnectorFormatCABLE
		powerTypeAC3PHASE := db.PowerTypeAC3PHASE

		cred := db.Credential{
			ID:          1,
			CountryCode: "FR",
			PartyID:     "GER",
		}

		location := db.Location{
			ID: 1,
		}

		evse := db.Evse{
			ID:         1,
			EvseID:     util.SqlNullString("DE-ABC-1234567"),
			Identifier: util.SqlNullString("DE*ABC*1234567"),
		}

		connectorDto := dto.ConnectorDto{
			Id:          util.NilString("1"),
			Standard:    &connectorTypeIEC62196T2,
			Format:      &connectorFormatCABLE,
			PowerType:   &powerTypeAC3PHASE,
			Voltage:     util.NilInt32(220),
			Amperage:    util.NilInt32(16),
			TariffID:    util.NilString("11"),
			LastUpdated: ocpitype.ParseOcpiTime("2015-03-16T10:10:02Z", nil),
		}

		connectorResolver.ReplaceConnector(ctx, cred, location, evse, *connectorDto.Id, &connectorDto)

		params, _ := mockRepository.GetCreateConnectorMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "1",
			"evseID": 1,
			"identifier": {"String": "DE*ABC*1234567*1", "Valid": true},
			"standard": "IEC_62196_T2",
			"format": "CABLE",
			"powerType": "AC_3_PHASE",
			"publish": true,
			"voltage": 220,
			"amperage": 16,
			"wattage": 10560,
			"tariffID": {"String": "11", "Valid": true},
			"termsAndConditions": {"String": "", "Valid": false},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))
	})

	t.Run("Update Connector", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		connectorResolver := connectorMocks.NewResolver(mockRepository)

		mockRepository.SetGetConnectorByEvseMockData(dbMocks.ConnectorMockData{
			Connector: db.Connector{
				Uid:         "1",
				EvseID:      1,
				Standard:    "IEC_62196_T2",
				Format:      "CABLE",
				PowerType:   "AC_3_PHASE",
				Voltage:     220,
				Amperage:    16,
				TariffID:    util.SqlNullString("11"),
				LastUpdated: *util.ParseTime("2015-03-16T10:10:02Z", nil),
			},
		})

		cred := db.Credential{
			ID:          1,
			CountryCode: "FR",
			PartyID:     "GER",
		}

		location := db.Location{
			ID: 1,
		}

		evse := db.Evse{
			ID: 1,
		}
		connectorDto := dto.ConnectorDto{
			TariffID: util.NilString("12"),
		}

		connectorResolver.ReplaceConnector(ctx, cred, location, evse, "1", &connectorDto)

		params, _ := mockRepository.GetUpdateConnectorByEvseMockData()
		paramsJson, _ := json.Marshal(params)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"uid": "1",
			"evseID": 1,
			"identifier": {"String": "", "Valid": false},
			"standard": "IEC_62196_T2",
			"format": "CABLE",
			"powerType": "AC_3_PHASE",
			"publish": true,
			"voltage": 220,
			"amperage": 16,
			"wattage": 10560,
			"tariffID": {"String": "12", "Valid": true},
			"termsAndConditions": {"String": "", "Valid": false},
			"lastUpdated": "2015-03-16T10:10:02Z"
		}`))
	})
}
