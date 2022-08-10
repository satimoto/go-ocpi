package mocks

import (
	"io"
	"net/http"
	"strings"
)

type MockResponseData struct {
	Response *http.Response
	Error    error
}

type MockHTTPRequester struct {
	requestData  []*http.Request
	responseData []MockResponseData
}

func (r *MockHTTPRequester) Do(req *http.Request) (*http.Response, error) {
	r.requestData = append(r.requestData, req)

	if len(r.responseData) == 0 {
		return nil, http.ErrBodyReadAfterClose
	}

	data := r.responseData[0]
	r.responseData = r.responseData[1:]
	return data.Response, data.Error
}

func (r *MockHTTPRequester) SetResponse(data MockResponseData) {
	r.responseData = append(r.responseData, data)
}

func (r *MockHTTPRequester) SetResponseWithBytes(statusCode int, bodyBytes string, err error) {
	readerCloser := io.NopCloser(strings.NewReader(bodyBytes))

	r.SetResponse(MockResponseData{
		Response: &http.Response{
			StatusCode: statusCode,
			Body:       readerCloser,
		},
		Error: err,
	})
}

func (r *MockHTTPRequester) GetRequest() *http.Request {
	if len(r.requestData) == 0 {
		return nil
	}

	data := r.requestData[0]
	r.requestData = r.requestData[1:]
	return data
}
