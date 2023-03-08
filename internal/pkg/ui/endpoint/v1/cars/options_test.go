/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package cars

import (
	appContract "HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/application/car/usecase"
	uiContract "HomeDepot/internal/pkg/ui/contract"
	"HomeDepot/internal/pkg/ui/endpoint"
	"net/http"
	"testing"
)

func TestOptions(t *testing.T) {
	sut := endpoint.NewSUT(
		func(r appContract.Repository) uiContract.Endpoint {
			return New(usecase.NewGetCars(r))
		},
		[]endpoint.Test{
			{
				"OPTIONS for " + root,
				endpoint.NoRepositoryInit,
				http.MethodOptions,
				root,
				endpoint.NoHeaderInit,
				endpoint.NoBodyInit,
				http.StatusOK,
				endpoint.OptionsHeaders([]string{http.MethodGet}),
				endpoint.EmptyBody,
				endpoint.NoRepositoryAssertion,
			},
		},
	)
	sut.Run(t)
}
