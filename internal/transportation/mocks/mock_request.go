package mocks

import (
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func NewOcpiService(requester *mocks.MockHTTPRequester) *transportation.OcpiService {
	return &transportation.OcpiService{
		HTTPRequester: requester,
	}
}
