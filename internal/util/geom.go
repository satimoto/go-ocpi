package util

import (
	"strconv"

	"github.com/paulmach/orb"
)

func NewPoint(latitude string, longitude string) (orb.Point, error) {
	var latitudeFloat, longitudeFloat float64
	var err error

	if latitudeFloat, err = strconv.ParseFloat(latitude, 64); err != nil {
		return orb.Point{}, err
	}

	if longitudeFloat, err = strconv.ParseFloat(longitude, 64); err != nil {
		return orb.Point{}, err
	}

	return orb.Point{latitudeFloat, longitudeFloat}, nil
}

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