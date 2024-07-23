package deise

type DeIseTariffDto struct {
	EvseID             string  `csv:"EVSE_ID"`
	ActivePriceSet     string  `csv:"ACTIVE_PRICE_SET"`
	ActiveKwhPrice     float64 `csv:"ACTIVE_KWH_PRICE"`
	ActiveMinutePrice  float64 `csv:"ACTIVE_MINUTE_PRICE"`
	ActiveMinuteDelay  int32   `csv:"ACTIVE_MINUTE_DELAY"`
	ActiveSessionPrice float64 `csv:"ACTIVE_SESSION_PRICE"`
	NextPriceSet       string  `csv:"-"`
	NextPriceTimestamp string  `csv:"-"`
	NextKwhPrice       string  `csv:"-"`
	NextMinutePrice    string  `csv:"-"`
	NextMinuteDelay    string  `csv:"-"`
	NextSessionPrice   string  `csv:"-"`
}
