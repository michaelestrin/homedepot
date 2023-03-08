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

func TestPost(t *testing.T) {
	modelWithoutID := *v1.TestModel
	modelWithoutID.ID = ""

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
				"POST with missing Content-Type",
				endpoint.NoRepositoryInit,
				http.MethodPost,
				root,
				endpoint.NoHeaderInit,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(v1.TestModel)))),
				http.StatusUnsupportedMediaType,
				endpoint.AllowOriginHeader,
				"{\"error\":\"missing or unsupported content type\"}",
				[]*car.Model{},
			},
			{
				"POST with invalid Content-Type",
				endpoint.NoRepositoryInit,
				http.MethodPost,
				root,
				endpoint.HeaderContentTypeInvalid,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(v1.TestModel)))),
				http.StatusUnsupportedMediaType,
				endpoint.AllowOriginHeader,
				"{\"error\":\"missing or unsupported content type\"}",
				[]*car.Model{},
			},
			{
				"POST with existing content",
				[]*car.Model{v1.TestModel},
				http.MethodPost,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(v1.TestModel)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"resource already exists\"}",
				[]*car.Model{v1.TestModel},
			},
			{
				"POST with new content missing ID",
				endpoint.NoRepositoryInit,
				http.MethodPost,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelWithoutID)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"consistency id is missing or empty\"}",
				[]*car.Model{},
			},
			{
				"POST with new content invalid year",
				endpoint.NoRepositoryInit,
				http.MethodPost,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelWithInvalidYear)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"consistency year must be \\u003e= 1900\"}",
				[]*car.Model{},
			},
			{
				"POST with new content negative mileage",
				endpoint.NoRepositoryInit,
				http.MethodPost,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelWithNegativeMileage)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"consistency mileage must be positive\"}",
				[]*car.Model{},
			},
			{
				"POST with new content negative price",
				endpoint.NoRepositoryInit,
				http.MethodPost,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(&modelWithNegativePrice)))),
				http.StatusBadRequest,
				endpoint.AllowOriginHeader,
				"{\"error\":\"consistency price must be positive\"}",
				[]*car.Model{},
			},
			{
				"POST with new valid content",
				endpoint.NoRepositoryInit,
				http.MethodPost,
				root,
				endpoint.HeaderContentTypeJSON,
				bytes.NewBuffer([]byte(v1.AsJSON(toDTO(v1.TestModel)))),
				http.StatusOK,
				endpoint.AllowOriginHeader,
				endpoint.EmptyBody,
				[]*car.Model{v1.TestModel},
			},
		},
	)
	sut.Run(t)
}
