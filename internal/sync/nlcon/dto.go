package nlcon

type NlConTariffDto struct {
	EvseID                      string   `json:"evseId"`
	TariffName                  string   `json:"tariffName"`
	StartTariff                 float64  `json:"startTariff"`
	EnergyTariff                float64  `json:"energyTariff"`
	ParkingGracePeriodInMinutes *int32   `json:"parkingGracePeriodInMinutes"`
	ParkingStepSize             *int64   `json:"parkingStepSize"`
	ParkingTariff               *float64 `json:"parkingTariff"`
}
