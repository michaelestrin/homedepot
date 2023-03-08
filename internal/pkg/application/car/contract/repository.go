/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package contract

import (
	"HomeDepot/internal/pkg/domain/car"
	"errors"
)

var (
	ErrResourceDoesNotExist  = errors.New("resource does not exist")
	ErrResourceAlreadyExists = errors.New("resource already exists")
)

type Repository interface {
	GetAll() ([]*car.Model, error)
	GetByID(id string) (*car.Model, error)
	Create(m *car.Model) error
	Update(m *car.Model) error
}
