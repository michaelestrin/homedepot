/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package usecase

import (
	"HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/domain/car"
)

type CreateCar struct {
	r contract.Repository
}

func NewCreateCar(r contract.Repository) *CreateCar {
	return &CreateCar{
		r: r,
	}
}

func (o *CreateCar) Do(m *car.Model) error {
	return o.r.Create(m)
}
