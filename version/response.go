package version

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"
)

type VersionPayload struct {
	Version *string `json:"version"`
	Url     *string `json:"url"`
}

func (r *VersionPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
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
