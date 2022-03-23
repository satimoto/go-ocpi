package versiondetail

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/satimoto/go-datastore/db"
)

var (
	version = "2.1.1"
)

type VersionDetailPayload struct {
	Version   string             `json:"version"`
	Endpoints []*EndpointPayload `json:"endpoints"`
}

func (r *VersionDetailPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type EndpointPayload struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
}

func NewCreateVersionEndpointParams(versionID int64, payload *EndpointPayload) db.CreateVersionEndpointParams {
	return db.CreateVersionEndpointParams{
		VersionID:  versionID,
		Identifier: payload.Identifier,
		Url:        payload.Url,
	}
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
	endpoints = append(endpoints, r.CreateEndpointPayload(ctx, apiDomain, "tariffs"))

	return &VersionDetailPayload{
		Version:   version,
		Endpoints: endpoints,
	}
}
