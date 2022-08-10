package evid_test

import (
	"testing"

	"github.com/satimoto/go-ocpi/pkg/evid"
)

func runEvIdTest(t *testing.T, value string, checksumResult string, validOk bool) {
	evId := evid.NewEvid(value)
	checksum := evId.ValueWithSeparator("-")

	if checksumResult != checksum {
		t.Errorf("Checksum error: %v, %v, %v", evId, checksum, checksumResult)
	}

	if evId.IsValid() != validOk {
		t.Errorf("Not valid: %v, %v", evId, checksum)
	}
}

func TestEvId(t *testing.T) {
	t.Run("Test DE*8AA*CA2B3C4D5", func(t *testing.T) {
		runEvIdTest(t, "DE*8AA*CA2B3C4D5", "DE-8AA-CA2B3C4D5-L", true)
	})

	t.Run("Test DE*8AA*CA2B3C4D5*L", func(t *testing.T) {
		runEvIdTest(t, "DE*8AA*CA2B3C4D5*L", "DE-8AA-CA2B3C4D5-L", true)
	})

	t.Run("Test DE-8AC-C12E56L89-2", func(t *testing.T) {
		runEvIdTest(t, "DE-8AC-C12E56L89-2", "DE-8AC-C12E56L89-2", false)
	})

	t.Run("Test DE-8AC-C12E56L89-4", func(t *testing.T) {
		runEvIdTest(t, "DE-8AC-C12E56L89-4", "DE-8AC-C12E56L89-4", true)
	})

	t.Run("Test DE-8AC-C12E56L89", func(t *testing.T) {
		runEvIdTest(t, "DE-8AC-C12E56L89", "DE-8AC-C12E56L89-4", true)
	})

	t.Run("Test BE-BEC-E041503002", func(t *testing.T) {
		runEvIdTest(t, "BE-BEC-E04150300-2", "BE-BEC-E04150300-2", false)
	})

	t.Run("NL*LMS*E17093261", func(t *testing.T) {
		runEvIdTest(t, "NL*LMS*E17093261", "NL-LMS-E17093261-K", true)
	})

	t.Run("Test DE*804*E1", func(t *testing.T) {
		runEvIdTest(t, "DE*804*E1", "", false)
	})

	t.Run("Get NL*LMS*E17093261 identifiers", func(t *testing.T) {
		evId := evid.NewEvid("NL*LMS*E17093261")
		countryCode := evId.GetCountryCode()
		partyID := evId.GetPartyID()

		if *countryCode != "NL" {
			t.Errorf("CountryCode error: %v, %v, %v", evId, countryCode, "NL")
		}

		if *partyID != "LMS" {
			t.Errorf("PartyID error: %v, %v, %v", evId, partyID, "LMS")
		}
	})

	t.Run("Get NL identifiers", func(t *testing.T) {
		evId := evid.NewEvid("NL")
		countryCode := evId.GetCountryCode()
		partyID := evId.GetPartyID()

		if countryCode != nil {
			t.Errorf("CountryCode error: %v, %v, %v", evId, countryCode, "NL")
		}

		if partyID != nil {
			t.Errorf("PartyID error: %v, %v, %v", evId, partyID, "LMS")
		}
	})
}
