package transportation

import (
	"net/http"
	"strconv"
)

type OcpiRequestHeader struct {
	Authorization *string
	ToCountryCode  *string
	ToPartyId      *string
}

func NewOcpiRequestHeader(token *string, countryCode *string, partyID *string) OcpiRequestHeader {
	return OcpiRequestHeader{
		Authorization: token,
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
