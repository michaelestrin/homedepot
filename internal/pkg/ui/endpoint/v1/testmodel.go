/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package v1

import (
	"HomeDepot/internal/pkg/domain/car"
	"encoding/json"
)

var TestModel = &car.Model{
	ID:        "id",
	Make:      "make",
	Model:     "TestModel",
	TrimLevel: "package",
	Color:     "color",
	Year:      1982,
	Category:  "category",
	Mileage:   1,
	Price:     2,
}

func AsJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
