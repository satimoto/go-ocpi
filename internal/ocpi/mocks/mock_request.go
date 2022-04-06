package mocks

import (
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func NewOCPIRequester(requester *mocks.MockHTTPRequester) *ocpi.OCPIRequester {
	return &ocpi.OCPIRequester{
		HTTPRequester: requester,
	}
}
