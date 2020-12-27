/*
 * telegram: @VasylNaumenko
 */

package rest

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"routes-api/pkg/api/v1/rest/handlers"
	"routes-api/pkg/config"

	"routes-api/pkg/log"
)

type Api struct {
	log log.Logger
	cfg config.Server

	hMeta handlers.Meta
	hOsrm handlers.Osrm
}

func New(
	cfg config.Server,
	log log.Logger,
	hStatus handlers.Meta,
	hOsrm handlers.Osrm) *Api {
	return &Api{
		cfg:   cfg,
		log:   log,
		hMeta: hStatus,
		hOsrm: hOsrm,
	}
}

func (a *Api) Run() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	a.registerRoutes(r)

	a.log.Infof("Listening %s", a.cfg.HTTP.ListenAddr)
	if err := http.ListenAndServe(a.cfg.HTTP.ListenAddr, r); err != nil {
		os.Exit(1)
	}
}
