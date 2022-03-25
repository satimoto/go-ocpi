package evid_test

import (
	"testing"

	"github.com/satimoto/go-ocpi-api/internal/evid"
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

	t.Run("Test DE*804*E1", func(t *testing.T) {
		runEvIdTest(t, "DE*804*E1", "", false)
	})
}
