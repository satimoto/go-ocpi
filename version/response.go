package version

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
)

type VersionPayload struct {
	Version *string `json:"version"`
	Url     *string `json:"url"`
}

func (r *VersionPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCreateVersionParams(credentialID int64, payload *VersionPayload) db.CreateVersionParams {
	return db.CreateVersionParams{
		CredentialID: credentialID,
		Version:      *payload.Version,
		Url:          *payload.Url,
	}
}

func (r *VersionResolver) CreateLocationPayload(ctx context.Context, apiDomain string, version string) *VersionPayload {
	url := fmt.Sprintf("%s/%s", apiDomain, version)

	return &VersionPayload{
		Version: &version,
		Url:     &url,
	}
}

func (r *VersionResolver) CreateVersionListPayload(ctx context.Context) []render.Renderer {
	apiDomain := os.Getenv("API_DOMAIN")

	list := []render.Renderer{}
	list = append(list, r.CreateLocationPayload(ctx, apiDomain, "2.1.1"))
	return list
}
