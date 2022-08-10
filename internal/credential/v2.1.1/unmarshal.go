package credential

import (
	"encoding/json"
	"io"
)


func (r *CredentialResolver) UnmarshalPullDto(body io.ReadCloser) (*OcpiCredentialDto, error) {
	response := OcpiCredentialDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
