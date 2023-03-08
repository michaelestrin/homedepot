/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package car

import (
	"HomeDepot/internal/pkg/application/car/usecase"
	"HomeDepot/internal/pkg/ui/contract"
	"HomeDepot/internal/pkg/ui/endpoint"
	"HomeDepot/internal/pkg/ui/endpoint/v1"
	"net/http"
)

const (
	path       = "/car"
	root       = v1.VersionPath + path
	rootForGet = root + "/"
)

type Endpoint struct {
	useCaseGetCar    *usecase.GetCar
	useCaseCreateCar *usecase.CreateCar
	useCaseUpdateCar *usecase.UpdateCar
}

func New(
	useCaseGetCar *usecase.GetCar,
	useCaseCreateCar *usecase.CreateCar,
	useCaseUpdateCar *usecase.UpdateCar,
) *Endpoint {
	return &Endpoint{
		useCaseGetCar:    useCaseGetCar,
		useCaseCreateCar: useCaseCreateCar,
		useCaseUpdateCar: useCaseUpdateCar,
	}
}

func (o *Endpoint) Register(s contract.Server) {
	s.Attach(
		root,
		func(rw http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				o.post(rw, r)
			case http.MethodPut:
				o.put(rw, r)
			case http.MethodOptions:
				endpoint.Options(rw, []string{http.MethodPost, http.MethodPut})
			default:
				endpoint.Respond(rw, nil, endpoint.ErrMethodNotAllowed)
			}
		},
	)
	s.Attach(
		rootForGet,
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
