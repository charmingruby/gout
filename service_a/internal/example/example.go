package example

import (
	"github.com/charmingruby/gout-service-a/internal/example/core/service"
	"github.com/charmingruby/gout-service-a/internal/example/database/dynamo"
	"github.com/charmingruby/gout-service-a/internal/example/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
)

func NewService() *service.Service {
	return service.New(service.Input{
		ExampleRepository: dynamo.NewExampleRepository(),
	})
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
