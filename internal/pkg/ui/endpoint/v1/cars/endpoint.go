/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package cars

import (
	"HomeDepot/internal/pkg/application/car/usecase"
	"HomeDepot/internal/pkg/ui/contract"
	"HomeDepot/internal/pkg/ui/endpoint"
	"HomeDepot/internal/pkg/ui/endpoint/v1"
	"net/http"
)

const (
	path = "/cars"
	root = v1.VersionPath + path
)

type Endpoint struct {
	useCaseGetCars *usecase.GetCars
}

func New(useCaseGetCars *usecase.GetCars) *Endpoint {
	return &Endpoint{
		useCaseGetCars: useCaseGetCars,
	}
}

func (o *Endpoint) Register(s contract.Server) {
	s.Attach(
		root,
		func(rw http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				o.get(rw, r)
			case http.MethodOptions:
				endpoint.Options(rw, []string{http.MethodGet})
			default:
				endpoint.Respond(rw, nil, endpoint.ErrMethodNotAllowed)
			}
		},
	)
}
