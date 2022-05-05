package transportation

import (
	"net/http"
	"strconv"
)

type OCPIRequestHeader struct {
	Authentication *string
	ToCountryCode *string
	ToPartyId *string
}

func NewOCPIRequestHeader(token *string, countryCode *string, partyID *string) OCPIRequestHeader {
	return OCPIRequestHeader{
		Authentication: token,
		ToCountryCode:  countryCode,
		ToPartyId:      partyID,
	}
}

func GetXLimitHeader(r *http.Response, defaultXLimit int) int {
	xLimitHeader := r.Header.Get("X-Limit")

	if len(xLimitHeader) > 0 {
		if xLimit, err := strconv.Atoi(xLimitHeader); err == nil {
			return xLimit
		}
	}

	return defaultXLimit
}