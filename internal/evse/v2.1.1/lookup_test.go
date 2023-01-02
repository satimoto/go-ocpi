package evse_test

import (
	"testing"

	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
)

func TestGetEvseIdentity(t *testing.T) {
	cases := []struct {
		name        string
		evseID      *string
		countryCode *string
		partyID     *string
	}{{
		"Test",
		util.NilString("BE-BEC-E041503008"),
		util.NilString("BE"),
		util.NilString("BEC"),
	}, {
		"Test",
		util.NilString(nil),
		util.NilString(nil),
		util.NilString(nil),
	}, {
		"FR*GIR",
		util.NilString(nil),
		util.NilString("FR"),
		util.NilString("GIR"),
	}, {
		"FR*GIR",
		util.NilString("FR-GIR-E04150300"),
		util.NilString("FR"),
		util.NilString("GIR"),
	}, {
		"FR-GIR",
		util.NilString("BE-BEC-E04150300"),
		util.NilString("FR"),
		util.NilString("GIR"),
	}, {
		"F-RGIR",
		util.NilString("EVB-P1415097*B2"),
		util.NilString(nil),
		util.NilString(nil),
	}, {
		"FR-GIR",
		util.NilString("EVB-P1415097*B2"),
		util.NilString("FR"),
		util.NilString("GIR"),
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			locationDto := dto.LocationDto{
				Operator: &coreDto.BusinessDetailDto{
					Name: tc.name,
				},
			}

			evseDto := dto.EvseDto{
				EvseID: tc.evseID,
			}

			countryCode, partyID := evse.GetEvseIdentity(&locationDto, &evseDto)

			if (countryCode == nil && tc.countryCode != nil) || (countryCode != nil && tc.countryCode == nil) || (countryCode != nil && tc.countryCode != nil && *countryCode != *tc.countryCode) ||
				(partyID == nil && tc.partyID != nil) || (partyID != nil && tc.partyID == nil) || (partyID != nil && tc.partyID != nil && *partyID != *tc.partyID) {
				t.Errorf("Value mismatch: %v %v expecting %v %v", util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""), util.DefaultString(tc.countryCode, ""), util.DefaultString(tc.partyID, ""))
			}
		})
	}
}
