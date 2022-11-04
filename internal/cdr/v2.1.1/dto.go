package cdr

import (
	"context"
	"log"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/metric"
)

func (r *CdrResolver) CreateCdrDto(ctx context.Context, cdr db.Cdr) *dto.CdrDto {
	response := dto.NewCdrDto(cdr)

	chargingPeriods, err := r.Repository.ListCdrChargingPeriods(ctx, cdr.ID)

	if err != nil {
		metrics.RecordError("OCPI223", "Error listing cdr charging periods", err)
		log.Printf("OCPI223: CdrID=%v", cdr.ID)
	} else {
		response.ChargingPeriods = r.ChargingPeriodResolver.CreateChargingPeriodListDto(ctx, chargingPeriods)
	}

	location, err := r.LocationResolver.Repository.GetLocation(ctx, cdr.LocationID)

	if err != nil {
		metrics.RecordError("OCPI224", "Error retrieving cdr location", err)
		log.Printf("OCPI224: LocationID=%v", cdr.LocationID)
	} else {
		response.Location = r.LocationResolver.CreateLocationDto(ctx, location)
	}

	if cdr.CalibrationID.Valid {
		calibration, err := r.CalibrationResolver.Repository.GetCalibration(ctx, cdr.CalibrationID.Int64)

		if err != nil {
			metrics.RecordError("OCPI225", "Error retrieving cdr calibration", err)
			log.Printf("OCPI225: CalibrationID=%v", cdr.CalibrationID)
		} else {
			response.SignedData = r.CalibrationResolver.CreateCalibrationDto(ctx, calibration)
		}
	}

	tariffs, err := r.TariffResolver.Repository.ListTariffsByCdr(ctx, util.SqlNullInt64(cdr.ID))

	if err != nil {
		metrics.RecordError("OCPI226", "Error listing cdr tariffs", err)
		log.Printf("OCPI226: CdrID=%v", cdr.ID)
	} else {
		response.Tariffs = r.TariffResolver.CreateTariffPushListDto(ctx, tariffs)
	}

	return response
}

func (r *CdrResolver) CreateCdrListDto(ctx context.Context, cdrs []db.Cdr) []render.Renderer {
	list := []render.Renderer{}

	for _, cdr := range cdrs {
		list = append(list, r.CreateCdrDto(ctx, cdr))
	}

	return list
}
