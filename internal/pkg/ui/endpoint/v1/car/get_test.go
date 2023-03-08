/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package car

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
			return New(usecase.NewGetCar(r), usecase.NewCreateCar(r), usecase.NewUpdateCar(r))
		},
		[]endpoint.Test{
			{
				"GET without ID in URL",
				endpoint.NoRepositoryInit,
				http.MethodGet,
				rootForGet,
				endpoint.NoHeaderInit,
				endpoint.NoBodyInit,
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"cannot parse id from url\"}",
				endpoint.NoRepositoryAssertion,
			},
			{
				"GET with valid ID",
				[]*car.Model{v1.TestModel},
				http.MethodGet,
				rootForGet + v1.TestModel.ID,
				endpoint.NoHeaderInit,
				endpoint.NoBodyInit,
				http.StatusOK,
				endpoint.AllowOriginHeader,
				v1.AsJSON(toDTO(v1.TestModel)),
				[]*car.Model{v1.TestModel},
			},
			{
				"GET with invalid ID",
				[]*car.Model{v1.TestModel},
				http.MethodGet,
				rootForGet + v1.TestModel.ID + "invalid",
				endpoint.NoHeaderInit,
				endpoint.NoBodyInit,
				http.StatusNotFound,
				endpoint.AllowOriginHeader,
				"{\"error\":\"resource does not exist\"}",
				[]*car.Model{v1.TestModel},
			},
		},
	)
	sut.Run(t)
}
