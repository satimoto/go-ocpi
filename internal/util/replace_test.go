package util_test

import (
	"testing"

	"github.com/satimoto/go-ocpi/internal/util"
)

func runTest(t *testing.T, str string, nth int, separator, expected string) {
	result := util.TrimFromNthSeparator(str, nth, separator)

	if result != expected {
		t.Errorf("Error: %v, %v, %v", str, result, expected)
	}

}

func TestTrimFromNthSeparator(t *testing.T) {
	t.Run("Test DE*8AA", func(t *testing.T) {
		runTest(t, "DE*8AA", 3, "*", "DE*8AA")
	})

	t.Run("Test DE*8AA*CA2B3C4D5", func(t *testing.T) {
		runTest(t, "DE*8AA*CA2B3C4D5", 3, "*", "DE*8AA*CA2B3C4D5")
	})

	t.Run("Test DE*8AA*CA2B3C4D5*L", func(t *testing.T) {
		runTest(t, "DE*8AA*CA2B3C4D5*L", 3, "*", "DE*8AA*CA2B3C4D5")
	})

	t.Run("Test DE*8AA*CA2B3C*5847", func(t *testing.T) {
		runTest(t, "DE*8AA*CA2B3C*5847", 3, "*", "DE*8AA*CA2B3C")
	})

	t.Run("Test DE*8AA*CA2B3C*5847", func(t *testing.T) {
		runTest(t, "DE*8AA*CA2B3C*5847", 4, "*", "DE*8AA*CA2B3C*5847")
	})
}
