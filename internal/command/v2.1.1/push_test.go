package command_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nsf/jsondiff"
	"github.com/satimoto/go-datastore/pkg/db"
	dbMocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1"
	commandMocks "github.com/satimoto/go-ocpi/internal/command/v2.1.1/mocks"
	transportationMocks "github.com/satimoto/go-ocpi/internal/transportation/mocks"
	"github.com/satimoto/go-ocpi/test/mocks"
)

func setupRoutes(commandResolver *command.CommandResolver) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/RESERVE_NOW/{command_id}", func(commandRouter chi.Router) {
		commandContextRouter := commandRouter.With(commandResolver.CommandReservationContext)
		commandContextRouter.Post("/", commandResolver.PostCommandReservationResponse)
	})

	router.Route("/START_SESSION/{command_id}", func(commandRouter chi.Router) {
		commandContextRouter := commandRouter.With(commandResolver.CommandStartContext)
		commandContextRouter.Post("/", commandResolver.PostCommandStartResponse)
	})

	router.Route("/STOP_SESSION/{command_id}", func(commandRouter chi.Router) {
		commandContextRouter := commandRouter.With(commandResolver.CommandStopContext)
		commandContextRouter.Post("/", commandResolver.PostCommandStopResponse)
	})

	router.Route("/UNLOCK_CONNECTOR/{command_id}", func(commandRouter chi.Router) {
		commandContextRouter := commandRouter.With(commandResolver.CommandUnlockContext)
		commandContextRouter.Post("/", commandResolver.PostCommandUnlockResponse)
	})

	return router
}

func TestCommandReservationRequest(t *testing.T) {
	t.Run("Invalid route", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/", nil)

		if err != nil {
			t.Fatal("Creating 'GET /' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusNotFound {
			t.Fatal("Returned ", responseRecorder.Code, " instead of ", http.StatusNotFound)
		}
	})

	t.Run("Invalid command", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodPost, "/RESERVE_NOW/1", bytes.NewBuffer([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /RESERVE_NOW/{command_id}' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 2003,
			"status_message": "Unknown resource"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Accept command reservation", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		cr := db.CommandReservation{
			ID:            1,
			Status:        db.CommandResponseTypeREQUESTED,
			TokenID:       1,
			ExpiryDate:    *util.ParseTime("2015-06-29T20:39:09Z", nil),
			ReservationID: 2,
			LocationID:    "LOC00001",
		}
		mockRepository.SetGetCommandReservationMockData(dbMocks.CommandReservationMockData{CommandReservation: cr, Error: nil})

		request, err := http.NewRequest(http.MethodPost, "/RESERVE_NOW/1", bytes.NewBuffer([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /RESERVE_NOW/{command_id}' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		commandParams, err := mockRepository.GetUpdateCommandReservationMockData()
		paramsJson, _ := json.Marshal(commandParams)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"id": 1,
			"status": "ACCEPTED",
			"expiryDate": "2015-06-29T20:39:09Z",
			"evseUid": {"String": "", "Valid": false}
		}`))

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})
}

func TestCommandStartRequest(t *testing.T) {
	t.Run("Invalid route", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/", nil)

		if err != nil {
			t.Fatal("Creating 'GET /' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusNotFound {
			t.Fatal("Returned ", responseRecorder.Code, " instead of ", http.StatusNotFound)
		}
	})

	t.Run("Invalid command", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodPost, "/START_SESSION/1", bytes.NewBuffer([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /START_SESSION/{command_id}' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 2003,
			"status_message": "Unknown resource"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Accept command start", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		cr := db.CommandStart{
			ID:         1,
			Status:     db.CommandResponseTypeREQUESTED,
			TokenID:    1,
			LocationID: "LOC00001",
		}
		mockRepository.SetGetCommandStartMockData(dbMocks.CommandStartMockData{CommandStart: cr, Error: nil})

		request, err := http.NewRequest(http.MethodPost, "/START_SESSION/1", bytes.NewBuffer([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /START_SESSION/{command_id}' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		commandParams, err := mockRepository.GetUpdateCommandStartMockData()
		paramsJson, _ := json.Marshal(commandParams)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"id": 1,
			"status": "ACCEPTED"
		}`))

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})
}

func TestCommandStopRequest(t *testing.T) {
	t.Run("Invalid route", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/", nil)

		if err != nil {
			t.Fatal("Creating 'GET /' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusNotFound {
			t.Fatal("Returned ", responseRecorder.Code, " instead of ", http.StatusNotFound)
		}
	})

	t.Run("Invalid command", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodPost, "/STOP_SESSION/1", bytes.NewBuffer([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /STOP_SESSION/{command_id}' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 2003,
			"status_message": "Unknown resource"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Accept command stop", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		cr := db.CommandStop{
			ID:        1,
			Status:    db.CommandResponseTypeREQUESTED,
			SessionID: "SESSION0001",
		}
		mockRepository.SetGetCommandStopMockData(dbMocks.CommandStopMockData{CommandStop: cr, Error: nil})

		request, err := http.NewRequest(http.MethodPost, "/STOP_SESSION/1", bytes.NewBuffer([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /STOP_SESSION/{command_id}' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		commandParams, err := mockRepository.GetUpdateCommandStopMockData()
		paramsJson, _ := json.Marshal(commandParams)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"id": 1,
			"status": "ACCEPTED"
		}`))

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})
}

func TestCommandUnlockRequest(t *testing.T) {
	t.Run("Invalid route", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/", nil)

		if err != nil {
			t.Fatal("Creating 'GET /' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusNotFound {
			t.Fatal("Returned ", responseRecorder.Code, " instead of ", http.StatusNotFound)
		}
	})

	t.Run("Invalid command", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodPost, "/UNLOCK_CONNECTOR/1", bytes.NewBuffer([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /UNLOCK_CONNECTOR/{command_id}' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 2003,
			"status_message": "Unknown resource"
		}`), jsondiff.SupersetMatch)
	})

	t.Run("Accept command unlock", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolverWithServices(mockRepository, transportationMocks.NewOcpiRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		cr := db.CommandUnlock{
			ID:          1,
			Status:      db.CommandResponseTypeREQUESTED,
			LocationID:  "LOC00001",
			EvseUid:     "EVSE0001",
			ConnectorID: "CONN0001",
		}
		mockRepository.SetGetCommandUnlockMockData(dbMocks.CommandUnlockMockData{CommandUnlock: cr, Error: nil})

		request, err := http.NewRequest(http.MethodPost, "/UNLOCK_CONNECTOR/1", bytes.NewBuffer([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /UNLOCK_CONNECTOR/{command_id}' request failed!")
		}

		commandRoutes.ServeHTTP(responseRecorder, request)

		commandParams, err := mockRepository.GetUpdateCommandUnlockMockData()
		paramsJson, _ := json.Marshal(commandParams)

		mocks.CompareJson(t, paramsJson, []byte(`{
			"id": 1,
			"status": "ACCEPTED"
		}`))

		mocks.CompareJsonWithDifference(t, responseRecorder.Body.Bytes(), []byte(`{
			"status_code": 1000,
			"status_message": "Success"
		}`), jsondiff.SupersetMatch)
	})
}
