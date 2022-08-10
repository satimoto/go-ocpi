package mocks

import (
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func NewOcpiRequester(requester *mocks.MockHTTPRequester) *transportation.OcpiRequester {
	return &transportation.OcpiRequester{
		HTTPRequester: requester,
	}
}
