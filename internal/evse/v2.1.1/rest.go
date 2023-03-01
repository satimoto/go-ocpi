package evse

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type EvseRestService struct {
	EvseResolver    *EvseResolver
	DispatchService *Dispatcher
}

func NewRestService(evseResolver *EvseResolver) *EvseRestService {
	maxWorkers := int(util.GetEnvInt32("EVSE_JOB_WORKERS", 5))
	dispatchService := NewDispatcher(maxWorkers)
	dispatchService.Start()

	return &EvseRestService{
		EvseResolver:    evseResolver,
		DispatchService: dispatchService,
	}
}

func (s *EvseRestService) GetEvse(rw http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	ctx := request.Context()
	evse := ctx.Value("evse").(db.Evse)
	dto := s.EvseResolver.CreateEvseDto(ctx, evse)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		metrics.RecordError("OCPI110", "Error rendering response", err)
		util.LogHttpRequest("OCPI110", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
	}
}

func (s *EvseRestService) UpdateEvse(rw http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	ctx := request.Context()
	cred := middleware.GetCredential(ctx)
	location := ctx.Value("location").(db.Location)
	uid := chi.URLParam(request, "evse_id")
	evseDto := dto.EvseDto{}

	if err := json.NewDecoder(request.Body).Decode(&evseDto); err != nil {
		metrics.RecordError("OCPI111", "Error unmarshaling request", err)
		util.LogHttpRequest("OCPI111", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	s.DispatchService.QueueJob(Job{
		EvseResolver: s.EvseResolver,
		Credential:   *cred,
		Location:     location,
		Uid:          uid,
		Dto:          evseDto,
	})

	render.Render(rw, request, transportation.OcpiSuccess(nil))
}

func (s *EvseRestService) Stop() {
	s.DispatchService.Stop()
}
