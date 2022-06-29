package evse

import "github.com/satimoto/go-ocpi/pkg/evid"

func GetEvseIdentity(evse *EvseDto) (*string, *string) {
	evseId := evid.NewEvid(*evse.EvseID)
	countryCode := evseId.GetCountryCode()
	partyID := evseId.GetPartyID()

	if countryCode != nil && partyID != nil {
		return countryCode, partyID
	}

	return nil, nil
}

func GetEvsesIdentity(evses []*EvseDto) (*string, *string) {
	for _, evseDto := range evses {
		if countryCode, partyID := GetEvseIdentity(evseDto); countryCode != nil && partyID != nil {
			return countryCode, partyID
		}
	}

	return nil, nil
}
