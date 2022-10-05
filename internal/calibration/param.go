package calibration

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateCalibrationParams(calibrationDto *coreDto.CalibrationDto) db.CreateCalibrationParams {
	return db.CreateCalibrationParams{
		EncodingMethod:        *calibrationDto.EncodingMethod,
		EncodingMethodVersion: util.SqlNullInt32(calibrationDto.EncodingMethodVersion),
		PublicKey:             util.SqlNullString(calibrationDto.PublicKey),
		Url:                   util.SqlNullString(calibrationDto.Url),
	}
}

func NewCreateCalibrationValueParams(id int64, calibrationValueDto *coreDto.CalibrationValueDto) db.CreateCalibrationValueParams {
	return db.CreateCalibrationValueParams{
		CalibrationID: id,
		Nature:        *calibrationValueDto.Nature,
		PlainData:     *calibrationValueDto.PlainData,
		SignedData:    *calibrationValueDto.SignedData,
	}
}
