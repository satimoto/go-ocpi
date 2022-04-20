package util

import "github.com/satimoto/go-datastore/db"

func CalculateWattage(powerType db.PowerType, voltage int32, amperage int32) int32 {
	if powerType == db.PowerTypeAC3PHASE {
		return 3 * 1 * amperage * voltage
	} else if powerType == db.PowerTypeAC1PHASE {
		return 1 * amperage * voltage
	}

	return amperage * voltage
}
