package evse_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	evseMocks "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateCapabilityListDto(t *testing.T) {
	ctx := context.Background()

	t.Run("Test", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		capabilities := []db.Capability{}
		capabilities = append(capabilities, db.Capability{Text: "Test1", Description: "Test"})
		capabilities = append(capabilities, db.Capability{Text: "Test2", Description: "Test"})
		capabilities = append(capabilities, db.Capability{Text: "Test3", Description: "Test"})

		response := evseResolver.CreateCapabilityListDto(ctx, capabilities)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`["Test1", "Test2", "Test3"]`))
	})
}

func TestCreateEvseDto(t *testing.T) {
	ctx := context.Background()

	t.Run("With Status Schedules", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		statusSchedules := []db.StatusSchedule{}
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T10:10:02Z"),
			PeriodEnd:   util.SqlNullTime(*util.ParseTime("2018-12-16T10:30:02Z")),
			Status:      db.EvseStatusBLOCKED,
		})
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T10:30:02Z"),
			PeriodEnd:   util.SqlNullTime(*util.ParseTime("2018-12-16T11:00:02Z")),
			Status:      db.EvseStatusCHARGING,
		})
		statusSchedules = append(statusSchedules, db.StatusSchedule{
			PeriodBegin: *util.ParseTime("2018-12-16T11:00:02Z"),
			Status:      db.EvseStatusAVAILABLE,
		})
		mockRepository.SetListStatusSchedulesMockData(dbMocks.StatusSchedulesMockData{StatusSchedules: statusSchedules, Error: nil})

		evse := db.Evse{
			Uid:               "3257",
			EvseID:            util.SqlNullString("BE-BEC-E041503002"),
			Status:            db.EvseStatusRESERVED,
			PhysicalReference: util.SqlNullString("2"),
			FloorLevel:        util.SqlNullString("-2"),
			LastUpdated:       *util.ParseTime("2015-06-29T20:39:09Z"),
		}

		response := evseResolver.CreateEvseDto(ctx, evse)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "3257",
			"evse_id": "BE-BEC-E041503002",
			"status": "RESERVED",
			"status_schedule": [{
				"period_begin": "2018-12-16T10:10:02Z",
				"period_end": "2018-12-16T10:30:02Z",
				"status": "BLOCKED"
			}, {
				"period_begin": "2018-12-16T10:30:02Z",
				"period_end": "2018-12-16T11:00:02Z",
				"status": "CHARGING"
			}, {
				"period_begin": "2018-12-16T11:00:02Z",
				"status": "AVAILABLE"
			}],
			"capabilities": [],
			"connectors": [],
			"physical_reference": "2",
			"floor_level": "-2",
			"directions": [],
			"parking_restrictions": [],
			"images": [],
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With Capabilities", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		capabilities := []db.Capability{}
		capabilities = append(capabilities, db.Capability{
			Text: "RESERVABLE",
		})
		mockRepository.SetListEvseCapabilitiesMockData(dbMocks.CapabilitiesMockData{Capabilities: capabilities, Error: nil})

		evse := db.Evse{
			Uid:         "3257",
			EvseID:      util.SqlNullString("BE-BEC-E041503002"),
			Status:      db.EvseStatusRESERVED,
			LastUpdated: *util.ParseTime("2015-06-29T20:39:09Z"),
		}

		response := evseResolver.CreateEvseDto(ctx, evse)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "3257",
			"evse_id": "BE-BEC-E041503002",
			"status": "RESERVED",
			"status_schedule": [],
			"capabilities": ["RESERVABLE"],
			"connectors": [],
			"directions": [],
			"parking_restrictions": [],
			"images": [],
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With Connectors", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		connectors := []db.Connector{}
		connectors = append(connectors, db.Connector{
			Uid:         "1",
			Standard:    "IEC_62196_T2",
			Format:      "CABLE",
			PowerType:   "AC_3_PHASE",
			Voltage:     220,
			Amperage:    16,
			TariffID:    util.SqlNullString("11"),
			LastUpdated: *util.ParseTime("2015-03-16T10:10:02Z"),
		})
		connectors = append(connectors, db.Connector{
			Uid:         "2",
			Standard:    "IEC_62196_T2_COMBO",
			Format:      "SOCKET",
			PowerType:   "AC_1_PHASE",
			Voltage:     110,
			Amperage:    32,
			TariffID:    util.SqlNullString("9"),
			LastUpdated: *util.ParseTime("2015-03-18T08:12:01Z"),
		})
		mockRepository.SetListConnectorsMockData(dbMocks.ConnectorsMockData{Connectors: connectors, Error: nil})

		evse := db.Evse{
			Uid:               "3256",
			EvseID:            util.SqlNullString("BE-BEC-E041503001"),
			Status:            db.EvseStatusAVAILABLE,
			PhysicalReference: util.SqlNullString("1"),
			FloorLevel:        util.SqlNullString("-1"),
			LastUpdated:       *util.ParseTime("2015-06-28T08:12:01Z"),
		}

		response := evseResolver.CreateEvseDto(ctx, evse)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "3256",
			"evse_id": "BE-BEC-E041503001",
			"status": "AVAILABLE",
			"status_schedule": [],
			"capabilities": [],
			"connectors": [{
				"id": "1",
				"standard": "IEC_62196_T2",
				"format": "CABLE",
				"power_type": "AC_3_PHASE",
				"voltage": 220,
				"amperage": 16,
				"tariff_id": "11",
				"last_updated": "2015-03-16T10:10:02Z"
			}, {
				"id": "2",
				"standard": "IEC_62196_T2_COMBO",
				"format": "SOCKET",
				"power_type": "AC_1_PHASE",
				"voltage": 110,
				"amperage": 32,
				"tariff_id": "9",
				"last_updated": "2015-03-18T08:12:01Z"
			}],
			"physical_reference": "1",
			"floor_level": "-1",
			"directions": [],
			"parking_restrictions": [],
			"images": [],
			"last_updated": "2015-06-28T08:12:01Z"
		}`))
	})

	t.Run("With Images", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		evseResolver := evseMocks.NewResolver(mockRepository)

		images := []db.Image{}
		images = append(images, db.Image{
			Url:      "https://business.com/logo.png",
			Category: db.ImageCategoryOPERATOR,
			Type:     "png",
		})
		images = append(images, db.Image{
			Url:       "https://business.com/logo2.jpg",
			Thumbnail: util.SqlNullString("https://business.com/logo2-thumb.jpg"),
			Category:  db.ImageCategoryENTRANCE,
			Type:      "jpg",
			Width:     sql.NullInt32{Int32: 180},
			Height:    sql.NullInt32{Int32: 180},
		})
		mockRepository.SetListEvseImagesMockData(dbMocks.ImagesMockData{Images: images, Error: nil})

		evse := db.Evse{
			Uid:               "3256",
			EvseID:            util.SqlNullString("BE-BEC-E041503001"),
			Status:            db.EvseStatusAVAILABLE,
			PhysicalReference: util.SqlNullString("1"),
			FloorLevel:        util.SqlNullString("-1"),
			LastUpdated:       *util.ParseTime("2015-06-28T08:12:01Z"),
		}

		response := evseResolver.CreateEvseDto(ctx, evse)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"uid": "3256",
			"evse_id": "BE-BEC-E041503001",
			"status": "AVAILABLE",
			"status_schedule": [],
			"capabilities": [],
			"connectors": [],
			"physical_reference": "1",
			"floor_level": "-1",
			"directions": [],
			"parking_restrictions": [],
			"images": [{
				"url": "https://business.com/logo.png",
				"category": "OPERATOR",
				"type": "png"
			}, {
				"url": "https://business.com/logo2.jpg",
				"thumbnail": "https://business.com/logo2-thumb.jpg",
				"category": "ENTRANCE",
				"type": "jpg",
				"width": 180,
				"height": 180
			}],
			"last_updated": "2015-06-28T08:12:01Z"
		}`))
	})
}
