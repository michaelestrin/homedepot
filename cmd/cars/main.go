/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package main

import (
	"HomeDepot/internal/pkg/application/car/usecase"
	"HomeDepot/internal/pkg/infrastructure/car/repository"
	"HomeDepot/internal/pkg/ui/contract"
	"HomeDepot/internal/pkg/ui/endpoint/v1/car"
	"HomeDepot/internal/pkg/ui/endpoint/v1/cars"
	"HomeDepot/internal/pkg/ui/metrics"
	"HomeDepot/internal/pkg/ui/server"
)

func main() {
	r := repository.NewCar()

	s := server.New("localhost:8080", metrics.Requests)
	s.Run([]contract.Endpoint{
		car.New(usecase.NewGetCar(r), usecase.NewCreateCar(r), usecase.NewUpdateCar(r)),
		cars.New(usecase.NewGetCars(r)),
	})
}
