package transportation_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func TestOcpiResponse(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ocpiResponse := transportation.OcpiResponse{
			Timestamp:     transportation.NewOcpiTime(util.ParseTime("2015-06-29T20:39:09Z", nil)),
			StatusCode:    1000,
			StatusMessage: "Success",
		}

		marshalledOcpiResponse, err := json.Marshal(ocpiResponse)
		
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		mocks.CompareJson(t, marshalledOcpiResponse, []byte(`{
			"status_code": 1000,
			"status_message": "Success",
			"timestamp": "2015-06-29T20:39:09Z"
		}`))

		readerCloser := io.NopCloser(strings.NewReader(string(marshalledOcpiResponse)))
		unmarshalledOcpiResponse, err := transportation.UnmarshalResponseDto(readerCloser)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if ocpiResponse != *unmarshalledOcpiResponse {
			t.Errorf("Mismatch: %v", unmarshalledOcpiResponse)
		}
	})

	t.Run("Success", func(t *testing.T) {
		ocpiResponse := transportation.OcpiResponse{
			Timestamp:     transportation.NewOcpiTime(util.ParseTime("2015-06-29T20:39:09Z", nil)),
			StatusCode:    1000,
			StatusMessage: "Success",
		}

		marshalledOcpiResponse := []byte(`{
			"status_code": 1000,
			"status_message": "Success",
			"timestamp": "2015-06-29T20:39:09"
		}`)

		readerCloser := io.NopCloser(strings.NewReader(string(marshalledOcpiResponse)))
		unmarshalledOcpiResponse, err := transportation.UnmarshalResponseDto(readerCloser)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if ocpiResponse != *unmarshalledOcpiResponse {
			t.Errorf("Mismatch: %v", unmarshalledOcpiResponse)
		}
	})
}
