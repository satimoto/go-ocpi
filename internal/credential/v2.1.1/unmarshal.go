package credential

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)


func (r *CredentialResolver) UnmarshalPushDto(body io.ReadCloser) (*dto.CredentialDto, error) {
	credentialDto := dto.CredentialDto{}

	if err := json.NewDecoder(body).Decode(&credentialDto); err != nil {
		return nil, err
	}

	return &credentialDto, nil
}

func (r *CredentialResolver) UnmarshalPullDto(body io.ReadCloser) (*dto.OcpiCredentialDto, error) {
	response := dto.OcpiCredentialDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
