package mocks

import (
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func NewOCPIRequester(requester *mocks.MockHTTPRequester) *transportation.OCPIRequester {
	return &transportation.OCPIRequester{
		HTTPRequester: requester,
	}
}
