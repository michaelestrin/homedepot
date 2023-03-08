/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package usecase

import (
	"HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/domain/car"
)

type UpdateCar struct {
	r contract.Repository
}

func NewUpdateCar(r contract.Repository) *UpdateCar {
	return &UpdateCar{
		r: r,
	}
}

func (o *UpdateCar) Do(m *car.Model) error {
	return o.r.Update(m)
}
