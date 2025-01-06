package endpoint

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/charmingruby/gout-service-b/internal/example/core/service"
	"github.com/charmingruby/gout-service-b/internal/example/transport/rest/dto/response"
	"github.com/charmingruby/gout-service-b/internal/example/transport/rest/dto/request"
	"github.com/charmingruby/gout-service-b/internal/shared/custom_err/core_err"
	"github.com/charmingruby/gout-service-b/internal/shared/transport/rest"
)

func (e *Endpoint) makeGreetingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := rest.ParseRequest[request.GreetingRequest](r)
		if err != nil {
			rest.BadRequestErrorResponse(w, err.Error())
			return
		}

		res, err := e.service.Greeting(service.GreetingParams{
			Name: req.Name,
		})

		if err != nil {
			var notFoundErr *core_err.ResourceNotFoundErr
			if errors.As(err, &notFoundErr) {
				rest.NotFoundErrorResponse(w, err.Error())
				return
			}

			rest.InternalServerErrorResponse(w)
			return
		}

		rest.OKResponse(w, "", response.GreetingResponse{
			Greeting: fmt.Sprintf("Long time no see! `%s` was managed.", res.ID),
		})
	}
}
