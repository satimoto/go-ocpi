package tariff_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	tariffMocks "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestCreateTariffPushDto(t *testing.T) {
	ctx := context.Background()

	t.Run("Empty", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tariffResolver := tariffMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		loc := db.Tariff{}

		response := tariffResolver.CreateTariffPushDto(ctx, loc)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "",
			"currency": "",
			"elements": [],
			"last_updated": "0001-01-01T00:00:00Z"
		}`))
	})

	t.Run("With alt text and element", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tariffResolver := tariffMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		tariffAltTexts := []db.DisplayText{}
		tariffAltTexts = append(tariffAltTexts, db.DisplayText{
			Language: "en",
			Text:     "2 euro p/hour",
		})
		mockRepository.SetListTariffAltTextsMockData(dbMocks.DisplayTextsMockData{DisplayTexts: tariffAltTexts, Error: nil})

		priceComponents := []db.PriceComponent{}
		priceComponents = append(priceComponents, db.PriceComponent{
			Type:     db.TariffDimensionTIME,
			Price:    2.5,
			StepSize: 300,
		})
		mockRepository.SetListPriceComponentsMockData(dbMocks.PriceComponentsMockData{PriceComponents: priceComponents, Error: nil})

		elements := []db.Element{}
		elements = append(elements, db.Element{})
		mockRepository.SetListElementsMockData(dbMocks.ElementsMockData{Elements: elements, Error: nil})

		tar := db.Tariff{
			Uid:          "TARIFF01",
			Currency:     "EUR",
			TariffAltUrl: sql.NullString{String: "https://ev-power.de/"},
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z"),
		}

		response := tariffResolver.CreateTariffPushDto(ctx, tar)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "TARIFF01",
			"currency": "EUR",
			"tariff_alt_text": [{
				"language": "en",
				"text": "2 euro p/hour"
			}],
			"tariff_alt_url": "https://ev-power.de/",
			"elements": [{
				"price_components": [{
					"type": "TIME",
					"price": 2.5,
					"step_size": 300
				}]
			}],
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With multiple elements", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tariffResolver := tariffMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		// 1
		priceComponents1 := []db.PriceComponent{}
		priceComponents1 = append(priceComponents1, db.PriceComponent{
			Type:     db.TariffDimensionFLAT,
			Price:    2.5,
			StepSize: 1,
		})
		mockRepository.SetListPriceComponentsMockData(dbMocks.PriceComponentsMockData{PriceComponents: priceComponents1, Error: nil})

		// 2
		priceComponents2 := []db.PriceComponent{}
		priceComponents2 = append(priceComponents2, db.PriceComponent{
			Type:     db.TariffDimensionTIME,
			Price:    1,
			StepSize: 900,
		})
		mockRepository.SetListPriceComponentsMockData(dbMocks.PriceComponentsMockData{PriceComponents: priceComponents2, Error: nil})

		restriction2 := db.Restriction{
			MaxPower: sql.NullFloat64{Float64: 32},
		}
		mockRepository.SetGetRestrictionMockData(dbMocks.RestrictionMockData{Restriction: restriction2, Error: nil})

		elements := []db.Element{}
		elements = append(elements, db.Element{})
		elements = append(elements, db.Element{
			RestrictionID: sql.NullInt64{Int64: 1, Valid: true},
		})
		mockRepository.SetListElementsMockData(dbMocks.ElementsMockData{Elements: elements, Error: nil})

		tar := db.Tariff{
			Uid:          "TARIFF01",
			Currency:     "EUR",
			TariffAltUrl: sql.NullString{String: "https://ev-power.de/"},
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z"),
		}

		response := tariffResolver.CreateTariffPushDto(ctx, tar)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "TARIFF01",
			"currency": "EUR",
			"tariff_alt_url": "https://ev-power.de/",
			"elements": [{
				"price_components": [{
					"type": "FLAT",
					"price": 2.5,
					"step_size": 1
				}]
			}, {
				"price_components": [{
					"type": "TIME",
					"price": 1,
					"step_size": 900
				}],
				"restrictions": {
					"max_power": 32
				}
			}],
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})

	t.Run("With energy mix", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		tariffResolver := tariffMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))

		energySources := []db.EnergySource{}
		energySources = append(energySources, db.EnergySource{
			Source:     db.EnergySourceCategoryGENERALGREEN,
			Percentage: 35.9,
		})
		energySources = append(energySources, db.EnergySource{
			Source:     db.EnergySourceCategoryGAS,
			Percentage: 6.3,
		})
		energySources = append(energySources, db.EnergySource{
			Source:     db.EnergySourceCategoryCOAL,
			Percentage: 33.2,
		})
		energySources = append(energySources, db.EnergySource{
			Source:     db.EnergySourceCategoryGENERALFOSSIL,
			Percentage: 2.9,
		})
		energySources = append(energySources, db.EnergySource{
			Source:     db.EnergySourceCategoryNUCLEAR,
			Percentage: 21.7,
		})
		mockRepository.SetListEnergySourcesMockData(dbMocks.EnergySourcesMockData{EnergySources: energySources, Error: nil})

		environImpact := []db.EnvironmentalImpact{}
		environImpact = append(environImpact, db.EnvironmentalImpact{
			Source: db.EnvironmentalImpactCategoryNUCLEARWASTE,
			Amount: 0.00006,
		})
		environImpact = append(environImpact, db.EnvironmentalImpact{
			Source: db.EnvironmentalImpactCategoryCARBONDIOXIDE,
			Amount: 372,
		})
		mockRepository.SetListEnvironmentalImpactsMockData(dbMocks.EnvironmentalImpactsMockData{EnvironmentalImpacts: environImpact, Error: nil})

		energyMix := db.EnergyMix{
			IsGreenEnergy:     false,
			SupplierName:      sql.NullString{String: "E.ON Energy Deutschland"},
			EnergyProductName: sql.NullString{String: "E.ON DirektStrom eco"},
		}
		mockRepository.SetGetEnergyMixMockData(dbMocks.EnergyMixMockData{EnergyMix: energyMix, Error: nil})

		tar := db.Tariff{
			Uid:          "TARIFF01",
			Currency:     "EUR",
			TariffAltUrl: sql.NullString{String: "https://ev-power.de/"},
			EnergyMixID:  sql.NullInt64{Int64: 1, Valid: true},
			LastUpdated:  *util.ParseTime("2015-06-29T20:39:09Z"),
		}

		response := tariffResolver.CreateTariffPushDto(ctx, tar)
		responseJson, _ := json.Marshal(response)

		mocks.CompareJson(t, responseJson, []byte(`{
			"id": "TARIFF01",
			"currency": "EUR",
			"tariff_alt_url": "https://ev-power.de/",
			"elements": [],
			"energy_mix": {
				"is_green_energy": false,
				"energy_sources": [{
					"source": "GENERAL_GREEN",
					"percentage": 35.9
				}, {
					"source": "GAS",
					"percentage": 6.3
				}, {
					"source": "COAL",
					"percentage": 33.2
				}, {
					"source": "GENERAL_FOSSIL",
					"percentage": 2.9
				}, {
					"source": "NUCLEAR",
					"percentage": 21.7
				}],
				"environ_impact": [{
					"source": "NUCLEAR_WASTE",
					"amount": 0.00006
				}, {
					"source": "CARBON_DIOXIDE",
					"amount": 372
				}],
				"supplier_name": "E.ON Energy Deutschland",
				"energy_product_name": "E.ON DirektStrom eco"
			},
			"last_updated": "2015-06-29T20:39:09Z"
		}`))
	})
}
