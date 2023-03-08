/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package car

import (
	"HomeDepot/internal/pkg/domain"
	"fmt"
	"strings"
)

const (
	MinYear     = 1900
	minIntValue = 0
)

type Model struct {
	ID        string
	Make      string
	Model     string
	TrimLevel string
	Color     string
	Year      int
	Category  string
	Mileage   int
	Price     int
}

func guardStringHasContent(key, value string) error {
	if len(strings.ReplaceAll(value, " ", "")) == 0 {
		return fmt.Errorf("%w %s is missing or empty", domain.ErrDomainConsistency, key)
	}
	return nil
}

func guardYear(key string, value int) error {
	if value < MinYear {
		return fmt.Errorf("%w %s must be >= %d", domain.ErrDomainConsistency, key, MinYear)
	}
	return nil
}

func guardPositiveInt(key string, value int) error {
	if value < minIntValue {
		return fmt.Errorf("%w %s must be positive", domain.ErrDomainConsistency, key)
	}
	return nil
}

func New(
	id string,
	make string,
	model string,
	trimLevel string,
	color string,
	year int,
	category string,
	mileage int,
	price int,
) (*Model, error) {

	if err := guardStringHasContent("id", id); err != nil {
		return nil, err
	}
	if err := guardYear("year", year); err != nil {
		return nil, err
	}
	if err := guardPositiveInt("mileage", mileage); err != nil {
		return nil, err
	}
	if err := guardPositiveInt("price", price); err != nil {
		return nil, err
	}

	return &Model{
		ID:        id,
		Make:      make,
		Model:     model,
		TrimLevel: trimLevel,
		Color:     color,
		Year:      year,
		Category:  category,
		Mileage:   mileage,
		Price:     price,
	}, nil
}
