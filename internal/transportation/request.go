package transportation

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPRequester interface {
	Do(req *http.Request) (*http.Response, error)
}

type OcpiRequester struct {
	HTTPRequester
}

func NewOcpiRequester() *OcpiRequester {
	return &OcpiRequester{
		HTTPRequester: &http.Client{},
	}
}

func (r *OcpiRequester) Do(method, url string, header OcpiRequestHeader, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	if body != nil {
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}

	if header.Authorization != nil && len(*header.Authorization) > 0 {
		request.Header.Set("Authorization", fmt.Sprintf("Token %s", *header.Authorization))
	}

	if header.ToCountryCode != nil && len(*header.ToCountryCode) > 0 {
		request.Header.Set("Ocpi-to-country-code", *header.ToCountryCode)
	}

	if header.ToPartyId != nil && len(*header.ToPartyId) > 0 {
		request.Header.Set("Ocpi-to-party-id", *header.ToPartyId)
	}

	response, err := r.HTTPRequester.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}
