package v1api

import "geomap/internal/app/logger"

type API struct {
	log *logger.Logger
}

func New(p Params) Result {
	// api.log.Debug("initialized package")
	return Result{API: API{log: p.Logger}}
}

func (a *API) Version() string {
	return "api_v1"
}