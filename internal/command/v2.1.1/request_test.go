package command_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/nsf/jsondiff"
	dbMocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-datastore/db"
	command "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1"
	commandMocks "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/test/mocks"
)

func setupRoutes(commandResolver *command.CommandResolver) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/reserve/{command_id}", func(commandRouter chi.Router) {
		commandContextRouter := commandRouter.With(commandResolver.CommandReservationContext)
		commandContextRouter.Post("/", commandResolver.PostCommandReservationResponse)
	})

	router.Route("/start/{command_id}", func(commandRouter chi.Router) {
		commandContextRouter := commandRouter.With(commandResolver.CommandStartContext)
		commandContextRouter.Post("/", commandResolver.PostCommandStartResponse)
	})

	router.Route("/stop/{command_id}", func(commandRouter chi.Router) {
		commandContextRouter := commandRouter.With(commandResolver.CommandStopContext)
		commandContextRouter.Post("/", commandResolver.PostCommandStopResponse)
	})

	router.Route("/unlock/{command_id}", func(commandRouter chi.Router) {
		commandContextRouter := commandRouter.With(commandResolver.CommandUnlockContext)
		commandContextRouter.Post("/", commandResolver.PostCommandUnlockResponse)
	})

	return router
}

func TestCommandReservationRequest(t *testing.T) {
	t.Run("Invalid route", func(t *testing.T) {
		mockRepository := dbMocks.NewMockRepositoryService()
		mockHTTPRequester := &mocks.MockHTTPRequester{}
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("GET", "/", nil)

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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("POST", "/reserve/1", bytes.NewReader([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /reserve/{command_id}' request failed!")
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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		cr := db.CommandReservation{
			ID: 1,
			Status: db.CommandResponseTypeREQUESTED,
			TokenID: 1,
			ExpiryDate: *util.ParseTime("2015-06-29T20:39:09Z"),
			ReservationID: 2,
			LocationID: "LOC00001",
		}
		mockRepository.SetGetCommandReservationMockData(dbMocks.CommandReservationMockData{CommandReservation: cr, Error: nil})

		request, err := http.NewRequest("POST", "/reserve/1", bytes.NewReader([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /reserve/{command_id}' request failed!")
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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("GET", "/", nil)

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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("POST", "/start/1", bytes.NewReader([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /start/{command_id}' request failed!")
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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		cr := db.CommandStart{
			ID: 1,
			Status: db.CommandResponseTypeREQUESTED,
			TokenID: 1,
			LocationID: "LOC00001",
		}
		mockRepository.SetGetCommandStartMockData(dbMocks.CommandStartMockData{CommandStart: cr, Error: nil})

		request, err := http.NewRequest("POST", "/start/1", bytes.NewReader([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /start/{command_id}' request failed!")
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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("GET", "/", nil)

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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("POST", "/stop/1", bytes.NewReader([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /stop/{command_id}' request failed!")
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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		cr := db.CommandStop{
			ID: 1,
			Status: db.CommandResponseTypeREQUESTED,
			SessionID: "SESSION0001",
		}
		mockRepository.SetGetCommandStopMockData(dbMocks.CommandStopMockData{CommandStop: cr, Error: nil})

		request, err := http.NewRequest("POST", "/stop/1", bytes.NewReader([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /stop/{command_id}' request failed!")
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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("GET", "/", nil)

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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		request, err := http.NewRequest("POST", "/unlock/1", bytes.NewReader([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /unlock/{command_id}' request failed!")
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
		commandResolver := commandMocks.NewResolver(mockRepository, mocks.NewOCPIRequester(mockHTTPRequester))
		commandRoutes := setupRoutes(commandResolver)
		responseRecorder := httptest.NewRecorder()

		cr := db.CommandUnlock{
			ID: 1,
			Status: db.CommandResponseTypeREQUESTED,
			LocationID: "LOC00001",
			EvseUid: "EVSE0001",
			ConnectorID: "CONN0001",
		}
		mockRepository.SetGetCommandUnlockMockData(dbMocks.CommandUnlockMockData{CommandUnlock: cr, Error: nil})

		request, err := http.NewRequest("POST", "/unlock/1", bytes.NewReader([]byte(`{
			"result": "ACCEPTED"
		}`)))

		if err != nil {
			t.Fatal("Creating 'POST /unlock/{command_id}' request failed!")
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
