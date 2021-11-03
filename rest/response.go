package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type OCPIResponse struct {
	Data          interface{} `json:"data,omitempty"`
	StatusCode    int16       `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Timestamp     time.Time   `json:"timestamp"`
}

func (response *OCPIResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	render.Status(request, 200)
	return nil
}

func OCPISuccess(data interface{}) render.Renderer {
	return &OCPIResponse{
		Data:          data,
		StatusCode:    1000,
		StatusMessage: "Success",
		Timestamp:     time.Now(),
	}
}

func OCPIClientError(data interface{}, message string) render.Renderer {
	return &OCPIResponse{
		Data:          data,
		StatusCode:    2000,
		StatusMessage: message,
		Timestamp:     time.Now(),
	}
}

func OCPIErrorMissingParameters(data interface{}) render.Renderer {
	return &OCPIResponse{
		Data:          data,
		StatusCode:    2001,
		StatusMessage: "Invalid or missing parameters",
		Timestamp:     time.Now(),
	}
}

func OCPIErrorNotEnoughInformation(data interface{}) render.Renderer {
	return &OCPIResponse{
		Data:          data,
		StatusCode:    2002,
		StatusMessage: "Not enough information",
		Timestamp:     time.Now(),
	}
}

func OCPIErrorUnknownResource(data interface{}) render.Renderer {
	return &OCPIResponse{
		Data:          data,
		StatusCode:    2003,
		StatusMessage: "Unknown resource",
		Timestamp:     time.Now(),
	}
}

func OCPIServerError(data interface{}, message string) render.Renderer {
	return &OCPIResponse{
		Data:          data,
		StatusCode:    3000,
		StatusMessage: message,
		Timestamp:     time.Now(),
	}
}
