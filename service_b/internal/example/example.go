package example

import (
	"github.com/charmingruby/gout-service-b/internal/example/core/service"
	"github.com/charmingruby/gout-service-b/internal/example/database/postgres"
	"github.com/charmingruby/gout-service-b/internal/example/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func NewService(db *sqlx.DB) (*service.Service, error) {
	exampleRepo, err := postgres.NewExampleRepository(db)
	if err != nil {
		return nil, err
	}
	
	return service.New(service.Input{
		ExampleRepository: exampleRepo,
	}), nil	
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
