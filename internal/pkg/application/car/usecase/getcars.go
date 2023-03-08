/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package usecase

import (
	"HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/domain/car"
)

type GetCars struct {
	r contract.Repository
}

func NewGetCars(r contract.Repository) *GetCars {
	return &GetCars{
		r: r,
	}
}

func (o *GetCars) Do() ([]*car.Model, error) {
	return o.r.GetAll()
}
