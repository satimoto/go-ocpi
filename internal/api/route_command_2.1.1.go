package api

import (
	"github.com/go-chi/chi/v5"
	command "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1"
)

func (rs *RouterService) mountCommands() *chi.Mux {
	commandResolver := command.NewResolver(rs.RepositoryService)
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
