package version

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
)

type VersionDto struct {
	Version *string `json:"version"`
	Url     *string `json:"url"`
}

func (r *VersionDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCreateVersionParams(credentialID int64, dto *VersionDto) db.CreateVersionParams {
	return db.CreateVersionParams{
		CredentialID: credentialID,
		Version:      *dto.Version,
		Url:          *dto.Url,
	}
}

func (r *VersionResolver) CreateLocationDto(ctx context.Context, apiDomain string, version string) *VersionDto {
	url := fmt.Sprintf("%s/%s", apiDomain, version)

	return &VersionDto{
		Version: &version,
		Url:     &url,
	}
}

func (r *VersionResolver) CreateVersionListDto(ctx context.Context) []render.Renderer {
	apiDomain := os.Getenv("API_DOMAIN")

	list := []render.Renderer{}
	list = append(list, r.CreateLocationDto(ctx, apiDomain, "2.1.1"))
	return list
}
