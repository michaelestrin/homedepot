/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package repository

import (
	"HomeDepot/internal/pkg/domain/car"
	"sync"
)

type Car struct {
	store map[string]*car.Model
	m     sync.Mutex
}

func NewCar() *Car {
	return &Car{
		store: make(map[string]*car.Model),
		m:     sync.Mutex{},
	}
}
