package util

import "strconv"

func NewGeomPoint(latitude string, longitude string) ([]float64, error) {
	var point []float64
	var value float64
	var err error

	if value, err = strconv.ParseFloat(latitude, 64); err != nil {
		return nil, err
	}

	point = append(point, value)

	if value, err = strconv.ParseFloat(longitude, 64); err != nil {
		return nil, err
	}

	point = append(point, value)

	return point, nil
}