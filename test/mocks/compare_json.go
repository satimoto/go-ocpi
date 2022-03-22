package mocks

import (
	"testing"

	"github.com/nsf/jsondiff"
)

var (
	diffOpts = jsondiff.DefaultJSONOptions()
)

func CompareJson(t *testing.T, actualBytes []byte, expectedBytes []byte) {
	diffRes, diff := jsondiff.Compare(expectedBytes, actualBytes, &diffOpts)

	if diffRes != jsondiff.FullMatch {
		t.Errorf("JSON mismatch: %v", diff)
		t.Logf("JSON: %v", string(actualBytes))
	}
}
