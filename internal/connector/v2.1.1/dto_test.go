package connector_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	connectorMocks "github.com/satimoto/go-ocpi/internal/connector/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestCreateConnectorDto(t *testing.T) {
	ctx := context.Background()

	t.Run("Test", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		connectorResolver := connectorMocks.NewResolver(mockRepository)

		connector := db.Connector{
			Uid:         "1",
			Standard:    "IEC_62196_T2",
			Format:      "CABLE",
			PowerType:   "AC_3_PHASE",
			Voltage:     220,
			Amperage:    16,
			TariffID:    util.SqlNullString("11"),
			LastUpdated: *util.ParseTime("2015-03-16T10:10:02Z", nil),
		}

		response := connectorResolver.CreateConnectorDto(ctx, connector)
		responseJson, _ := json.Marshal(response)

		connectorParams := db.CreateConnectorParams{
			Uid:         "1",
			Standard:    "IEC_62196_T2",
			Format:      "CABLE",
			PowerType:   "AC_3_PHASE",
			Voltage:     220,
			Amperage:    16,
			TariffID:    util.SqlNullString("11"),
			LastUpdated: *util.ParseTime("2015-03-16T10:10:02Z", nil),
		}

		connectorResolver.Repository.CreateConnector(ctx, connectorParams)

		c, err := mockRepository.GetCreateConnectorMockData()
		t.Log(c, err)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "1",
			"standard": "IEC_62196_T2",
			"format": "CABLE",
			"power_type": "AC_3_PHASE",
			"voltage": 220,
			"amperage": 16,
			"tariff_id": "11",
			"last_updated": "2015-03-16T10:10:02Z"
		}`))
	})
}
