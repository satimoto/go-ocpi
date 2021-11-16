package util

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPRequester interface {
	Do(req *http.Request) (*http.Response, error)
}

type OCPIRequester struct {
	HTTPRequester
}

type OCPIRequestHeader struct {
	Authentication *string
	ToCountryCode *string
	ToPartyId *string
}

func NewOCPIRequester() *OCPIRequester {
	return &OCPIRequester{
		HTTPRequester: &http.Client{},
	}
}

func (r *OCPIRequester) Do(method, url string, header OCPIRequestHeader, body io.Reader) (io.ReadCloser, error) {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	if len(*header.Authentication) > 0 {
		request.Header.Set("Authentication", fmt.Sprintf("Token %s", *header.Authentication))
	}

	if len(*header.ToCountryCode) > 0 {
		request.Header.Set("OCPI-to-country-code", *header.ToCountryCode)
	}

	if len(*header.ToPartyId) > 0 {
		request.Header.Set("OCPI-to-party-id", *header.ToPartyId)
	}

	response, err := r.HTTPRequester.Do(request)

	if err != nil {
		return nil, err
	}
	
	return response.Body, nil
}
