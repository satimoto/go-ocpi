package geolocation

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/geom"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *GeoLocationResolver) ReplaceGeoLocation(ctx context.Context, id *sql.NullInt64, geoLocationDto *coreDto.GeoLocationDto) *geom.Geometry4326 {
	if geoLocationDto != nil {
		if id.Valid {
			updateGeoLocationParams := NewUpdateGeoLocationParams(id.Int64, geoLocationDto)
			_, err := r.Repository.UpdateGeoLocation(ctx, updateGeoLocationParams)

			if err != nil {
				metrics.RecordError("OCPI114", "Error updating geolocation", err)
				log.Printf("OCPI114: Params=%#v", updateGeoLocationParams)
				return nil
			}
		} else {
			createGeoLocationParams := NewCreateGeoLocationParams(geoLocationDto)
			geoLocation, err := r.Repository.CreateGeoLocation(ctx, createGeoLocationParams)

			if err != nil {
				metrics.RecordError("OCPI115", "Error creating geolocation", err)
				log.Printf("OCPI115: Params=%#v", createGeoLocationParams)
				return nil
			}

			id.Scan(geoLocation.ID)
		}

		point, err := geom.NewPoint(geoLocationDto.Longitude.String(), geoLocationDto.Latitude.String())

		if err != nil {
			metrics.RecordError("OCPI116", "Error creating geom point", err)
			log.Printf("OCPI116: Dto=%#v", geoLocationDto)
			return nil
		}

		return &geom.Geometry4326{
			Coordinates: point,
			Type:        point.GeoJSONType(),
		}
	}

	return nil
}
