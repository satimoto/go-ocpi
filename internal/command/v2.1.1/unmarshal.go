package command

import (
	"encoding/json"
	"io"
)

func (r *CommandResolver) UnmarshalCommandResponsePayload(body io.ReadCloser) (*CommandResponsePayload, error) {
	if body != nil {
		payload := CommandResponsePayload{}

		if err := json.NewDecoder(body).Decode(&payload); err != nil {
			return nil, err
		}

		return &payload, nil
	}

	return nil, nil
}
