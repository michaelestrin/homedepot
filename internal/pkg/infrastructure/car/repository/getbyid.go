/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package repository

import (
	"HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/domain/car"
)

func (o *Car) GetByID(id string) (*car.Model, error) {
	o.m.Lock()
	defer o.m.Unlock()

	c, ok := o.store[id]
	if !ok {
		return nil, contract.ErrResourceDoesNotExist
	}
	return c, nil
}
