package credential

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)


func (r *CredentialResolver) UnmarshalPullDto(body io.ReadCloser) (*dto.OcpiCredentialDto, error) {
	response := dto.OcpiCredentialDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
