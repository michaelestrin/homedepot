/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package cars

import (
	appContract "HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/application/car/usecase"
	"HomeDepot/internal/pkg/domain/car"
	uiContract "HomeDepot/internal/pkg/ui/contract"
	"HomeDepot/internal/pkg/ui/endpoint"
	v1 "HomeDepot/internal/pkg/ui/endpoint/v1"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	sut := endpoint.NewSUT(
		func(r appContract.Repository) uiContract.Endpoint {
			return New(usecase.NewGetCars(r))
		},
		[]endpoint.Test{
			{
				"GET with empty repo",
				endpoint.NoRepositoryInit,
				http.MethodGet,
				root,
				endpoint.NoHeaderInit,
				endpoint.NoBodyInit,
				http.StatusOK,
				endpoint.AllowOriginHeader,
				v1.AsJSON(toDTO([]*car.Model{})),
				nil,
			},
			{
				"GET with existing",
				[]*car.Model{v1.TestModel},
				http.MethodGet,
				root,
				endpoint.NoHeaderInit,
				endpoint.NoBodyInit,
				http.StatusOK,
				endpoint.AllowOriginHeader,
				v1.AsJSON(toDTO([]*car.Model{v1.TestModel})),
				nil,
			},
		},
	)
	sut.Run(t)
}
