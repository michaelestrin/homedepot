/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package usecase

import (
	"HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/domain/car"
)

type GetCar struct {
	r contract.Repository
}

func NewGetCar(r contract.Repository) *GetCar {
	return &GetCar{
		r: r,
	}
}

func (o *GetCar) Do(id string) (*car.Model, error) {
	return o.r.GetByID(id)
}
