/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package cars

import (
	"HomeDepot/internal/pkg/domain/car"
)

type dto struct {
	ID        string `json:"id"`
	Make      string `json:"make"`
	Model     string `json:"model"`
	TrimLevel string `json:"package"`
	Color     string `json:"color"`
	Year      int    `json:"year"`
	Category  string `json:"category"`
	Mileage   int    `json:"mileage"`
	Price     int    `json:"price"`
}

func toDTO(m []*car.Model) []*dto {
	cars := make([]*dto, 0)
	for _, c := range m {
		cars = append(
			cars,
			&dto{
				ID:        c.ID,
				Make:      c.Make,
				Model:     c.Model,
				TrimLevel: c.TrimLevel,
				Color:     c.Color,
				Year:      c.Year,
				Category:  c.Category,
				Mileage:   c.Mileage,
				Price:     c.Price,
			},
		)
	}
	return cars
}
