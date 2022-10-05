package evse

import (
	"regexp"
	"strings"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/util"
	"github.com/satimoto/go-ocpi/pkg/evid"
)

func GetEvseIdentifier(evseDto *dto.EvseDto) *string {
	if evseDto.EvseID != nil {
		identifier := util.TrimFromNthSeparator(strings.ReplaceAll(*evseDto.EvseID, "-", "*"), 3, "*")

		return &identifier
	}
	
	return nil
}

func GetEvseIdentity(locationDto *dto.LocationDto, evseDto *dto.EvseDto) (*string, *string) {
	if locationDto.Operator != nil {
		operatorName := locationDto.Operator.Name
		found, err := regexp.MatchString("^[A-Za-z]{2}[-*][A-Za-z0-9]{3}$", operatorName)
		
		if found && err == nil {
			countryCode := operatorName[:2]
			partyID := operatorName[3:6]

			return &countryCode, &partyID
		}
	}

	if evseDto.EvseID != nil {
		evseId := evid.NewEvid(*evseDto.EvseID)

		if evseId.IsValid() {
			countryCode := evseId.GetCountryCode()
			partyID := evseId.GetPartyID()	

			if countryCode != nil && partyID != nil {
				return countryCode, partyID
			}
		}
	}

	return nil, nil
}

func GetEvsesIdentity(locationDto *dto.LocationDto, evses []*dto.EvseDto) (*string, *string) {
	for _, evseDto := range evses {
		if countryCode, partyID := GetEvseIdentity(locationDto, evseDto); countryCode != nil && partyID != nil {
			return countryCode, partyID
		}
	}

	return nil, nil
}
