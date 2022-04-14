package calibration

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func NewCreateCalibrationParams(dto *CalibrationDto) db.CreateCalibrationParams {
	return db.CreateCalibrationParams{
		EncodingMethod:        *dto.EncodingMethod,
		EncodingMethodVersion: util.SqlNullInt32(dto.EncodingMethodVersion),
		PublicKey:             util.SqlNullString(dto.PublicKey),
		Url:                   util.SqlNullString(dto.Url),
	}
}

func NewCreateCalibrationValueParams(id int64, dto *CalibrationValueDto) db.CreateCalibrationValueParams {
	return db.CreateCalibrationValueParams{
		CalibrationID: id,
		Nature:        *dto.Nature,
		PlainData:     *dto.PlainData,
		SignedData:    *dto.SignedData,
	}
}
