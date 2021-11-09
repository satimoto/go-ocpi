package version

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

var (
	version   = "2.1.1"
)

type VersionDetailPayload struct {
	Version   string             `json:"version"`
	Endpoints []*EndpointPayload `json:"endpoints"`
}

type EndpointPayload struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
}

func (r *VersionDetailPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (r *VersionDetailResolver) CreateEndpointPayload(ctx context.Context, apiDomain string, identifier string) *EndpointPayload {
	return &EndpointPayload{
		Identifier: identifier,
		Url:        fmt.Sprintf("%s/%s/%s", apiDomain, version, identifier),
	}
}

func (r *VersionDetailResolver) CreateVersionDetailPayload(ctx context.Context) *VersionDetailPayload {
	apiDomain := os.Getenv("API_DOMAIN")

	var endpoints []*EndpointPayload
	endpoints = append(endpoints, r.CreateEndpointPayload(ctx, apiDomain, "locations"))

	return &VersionDetailPayload{
		Version:   version,
		Endpoints: endpoints,
	}
}
