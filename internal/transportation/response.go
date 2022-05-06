package transportation

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

const (
	STATUS_CODE_OK                 = 1000
	STATUS_CODE_CLIENT_ERROR       = 2000
	STATUS_CODE_MISSING_PARAMS     = 2001
	STATUS_CODE_NOT_ENOUGH_INFO    = 2002
	STATUS_CODE_UNKNOWN_RESOURCE   = 2003
	STATUS_CODE_SERVER_ERROR       = 3000
	STATUS_CODE_REGISTRATION_ERROR = 3001
	STATUS_CODE_UNSUPPRTED_VERSION = 3002
	STATUS_CODE_MISSING_ENDPOINTS  = 3003
)

type OcpiResponse struct {
	Data          interface{} `json:"data,omitempty"`
	StatusCode    int16       `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Timestamp     time.Time   `json:"timestamp"`
}

func (response *OcpiResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	render.Status(request, 200)
	return nil
}

func (response *OcpiResponse) Error() string {
	return response.StatusMessage
}

func OcpiSuccess(data interface{}) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_OK,
		StatusMessage: "Success",
		Timestamp:     time.Now(),
	}
}

func OcpiClientError(data interface{}, message string) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_CLIENT_ERROR,
		StatusMessage: message,
		Timestamp:     time.Now(),
	}
}

func OcpiErrorMissingParameters(data interface{}) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_MISSING_PARAMS,
		StatusMessage: "Invalid or missing parameters",
		Timestamp:     time.Now(),
	}
}

func OcpiErrorNotEnoughInformation(data interface{}) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_NOT_ENOUGH_INFO,
		StatusMessage: "Not enough information",
		Timestamp:     time.Now(),
	}
}

func OcpiErrorUnknownResource(data interface{}) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_UNKNOWN_RESOURCE,
		StatusMessage: "Unknown resource",
		Timestamp:     time.Now(),
	}
}

func OcpiServerError(data interface{}, message string) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_SERVER_ERROR,
		StatusMessage: message,
		Timestamp:     time.Now(),
	}
}

func OcpiRegistrationError(data interface{}) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_REGISTRATION_ERROR,
		StatusMessage: "Registration error",
		Timestamp:     time.Now(),
	}
}

func OcpiUnsupportedVersion(data interface{}) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_UNSUPPRTED_VERSION,
		StatusMessage: "Unsupported version",
		Timestamp:     time.Now(),
	}
}

func OcpiMissingEndpoints(data interface{}) *OcpiResponse {
	return &OcpiResponse{
		Data:          data,
		StatusCode:    STATUS_CODE_MISSING_ENDPOINTS,
		StatusMessage: "Missing endpoints",
		Timestamp:     time.Now(),
	}
}
