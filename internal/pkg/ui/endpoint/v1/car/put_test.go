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
	"bytes"
	"net/http"
	"testing"
)

func TestPut(t *testing.T) {
	modelWithoutID := *v1.TestModel
	modelWithoutID.ID = ""

	modelModified := *v1.TestModel
	modelModified.Make += "modified"

	modelWithInvalidYear := *v1.TestModel
	modelWithInvalidYear.Year = car.MinYear - 1

	modelWithNegativeMileage := *v1.TestModel
	modelWithNegativeMileage.Mileage = -1

	modelWithNegativePrice := *v1.TestModel
	modelWithNegativePrice.Price = -1

	sut := endpoint.NewSUT(
		func(r appContract.Repository) uiContract.Endpoint {
			return New(usecase.NewGetCar(r), usecase.NewCreateCar(r), usecase.NewUpdateCar(r))
		},
		[]endpoint.Test{
			{
				"PUT with missing Content-Type",
				endpoint.NoRepositoryInit,
				http.MethodPut,
				root,
				endpoint.NoHeaderInit,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(v1.TestModel)))),
				http.StatusUnsupportedMediaType,
				endpoint.AllowOriginHeader,
				"{\"error\":\"missing or unsupported content type\"}",
				[]*car.Model{},
			},
			{
				"PUT with invalid Content-Type",
				endpoint.NoRepositoryInit,
				http.MethodPut,
				root,
				endpoint.HeaderContentTypeInvalid,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(v1.TestModel)))),
				http.StatusUnsupportedMediaType,
				endpoint.AllowOriginHeader,
				"{\"error\":\"missing or unsupported content type\"}",
				[]*car.Model{},
			},
			{
				"PUT with new content missing ID",
				endpoint.NoRepositoryInit,
				http.MethodPut,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelWithoutID)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"consistency id is missing or empty\"}",
				[]*car.Model{},
			},
			{
				"PUT with new content invalid year",
				endpoint.NoRepositoryInit,
				http.MethodPut,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelWithInvalidYear)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"consistency year must be \\u003e= 1900\"}",
				[]*car.Model{},
			},
			{
				"PUT with new content negative mileage",
				endpoint.NoRepositoryInit,
				http.MethodPut,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelWithNegativeMileage)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"consistency mileage must be positive\"}",
				[]*car.Model{},
			},
			{
				"PUT with new content negative price",
				endpoint.NoRepositoryInit,
				http.MethodPut,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelWithNegativePrice)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"consistency price must be positive\"}",
				[]*car.Model{},
			},
			{
				"PUT without existing ID",
				endpoint.NoRepositoryInit,
				http.MethodPut,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(v1.TestModel)))),
				http.StatusNotFound,
				endpoint.AllowOriginHeader,
				"{\"error\":\"resource does not exist\"}",
				[]*car.Model{},
			},
			{
				"PUT with existing content",
				[]*car.Model{v1.TestModel},
				http.MethodPut,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelModified)))),
				http.StatusOK,
				endpoint.AllowOriginHeader,
				endpoint.EmptyBody,
				[]*car.Model{&modelModified},
			},
		},
	)
	sut.Run(t)
}
