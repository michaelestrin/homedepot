/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package repository

import (
	"HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/domain/car"
)

func (o *Car) Update(m *car.Model) error {
	o.m.Lock()
	defer o.m.Unlock()

	if _, ok := o.store[m.ID]; !ok {
		return contract.ErrResourceDoesNotExist
	}

	o.store[m.ID] = m
	return nil
}
