package version

import (
	"context"
	"fmt"
	"os"

	"github.com/go-chi/render"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *VersionResolver) CreateLocationDto(ctx context.Context, apiDomain string, version string) *coreDto.VersionDto {
	url := fmt.Sprintf("%s/%s", apiDomain, version)

	return &coreDto.VersionDto{
		Version: &version,
		Url:     &url,
	}
}

func (r *VersionResolver) CreateVersionListDto(ctx context.Context) []render.Renderer {
	apiDomain := os.Getenv("API_DOMAIN")

	list := []render.Renderer{}
	list = append(list, r.CreateLocationDto(ctx, apiDomain, VERSION_2_1_1))
	return list
}
