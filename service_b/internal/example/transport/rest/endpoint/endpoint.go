package endpoint

import (
	"github.com/charmingruby/gout-service-b/internal/example/core/service"
	"github.com/go-chi/chi/v5"
)

type Endpoint struct {
	router *chi.Mux
	service *service.Service
}

func New(r *chi.Mux, service *service.Service) *Endpoint {
	return &Endpoint{
		router: r,
		service: service,
	}
}

func (e *Endpoint) Register() {
	e.router.Post("/example/greeting", e.makeGreetingHandler())
}
