package businessdetail_test

import (
	"encoding/json"
	"testing"

	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func TestBusinessDetailUnmarshal(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		businessDetailDto := coreDto.BusinessDetailDto{}
		response := []byte(`{
			"name": ""
		}`)

		json.Unmarshal([]byte(`{}`), &businessDetailDto)
		responseJson, _ := json.Marshal(businessDetailDto)

		mocks.CompareJson(t, responseJson, response)
	})

	t.Run("Name only", func(t *testing.T) {
		businessDetailDto := coreDto.BusinessDetailDto{}
		request := []byte(`{
			"name": "Business Name"
		}`)

		json.Unmarshal(request, &businessDetailDto)
		responseJson, _ := json.Marshal(businessDetailDto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("Name and Website", func(t *testing.T) {
		businessDetailDto := coreDto.BusinessDetailDto{}
		request := []byte(`{
			"name": "Business Name",
			"website": "https://business.com"
		}`)

		json.Unmarshal(request, &businessDetailDto)
		responseJson, _ := json.Marshal(businessDetailDto)

		mocks.CompareJson(t, responseJson, request)
	})

	t.Run("Name, Website and Logo", func(t *testing.T) {
		businessDetailDto := coreDto.BusinessDetailDto{}
		request := []byte(`{
			"name": "Business Name",
			"website": "https://business.com",
			"logo": {
				"url": "https://business.com/logo.png",
				"category": "OPERATOR",
				"type": "png"
			}
		}`)

		json.Unmarshal(request, &businessDetailDto)
		responseJson, _ := json.Marshal(businessDetailDto)

		mocks.CompareJson(t, responseJson, request)
	})
}
