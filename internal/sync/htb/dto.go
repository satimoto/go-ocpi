package htb

type PriceInformationDto struct {
	EvseID     string        `json:"evse_id"`
	Conditions *ConditionsDto `json:"conditions"`
}

type ConditionsDto struct {
	Rate *string `json:"rate"`
}
