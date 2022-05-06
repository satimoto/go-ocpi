package mocks

import (
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func NewOcpiRequester(requester *mocks.MockHTTPRequester) *transportation.OcpiRequester {
	return &transportation.OcpiRequester{
		HTTPRequester: requester,
	}
}
