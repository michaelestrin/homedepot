/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package repository

import (
	"HomeDepot/internal/pkg/domain/car"
)

func (o *Car) GetAll() ([]*car.Model, error) {
	o.m.Lock()
	defer o.m.Unlock()

	cars := make([]*car.Model, 0)
	for _, c := range o.store {
		cars = append(cars, c)
	}
	return cars, nil
}
