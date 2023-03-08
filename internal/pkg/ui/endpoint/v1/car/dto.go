/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package car

import (
	"HomeDepot/internal/pkg/domain/car"
	"HomeDepot/internal/pkg/ui/endpoint"
	"encoding/json"
	"net/http"
	"strings"
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

func toDTO(m *car.Model) *dto {
	if m == nil {
		return nil
	}

	return &dto{
		ID:        m.ID,
		Make:      m.Make,
		Model:     m.Model,
		TrimLevel: m.TrimLevel,
		Color:     m.Color,
		Year:      m.Year,
		Category:  m.Category,
		Mileage:   m.Mileage,
		Price:     m.Price,
	}
}

func supportedContentType(r *http.Request) bool {
	ct, ok := r.Header[endpoint.HeaderContentTypeKey]
	if !ok {
		return false
	}
	for _, h := range ct {
		if strings.Contains(h, endpoint.HeaderContentTypeValueJSON) {
			return true
		}
	}
	return false
}

func toModel(r *http.Request) (*car.Model, error) {
	if r.ContentLength == 0 {
		return nil, endpoint.ErrContentLengthIsZero
	}

	if !supportedContentType(r) {
		return nil, endpoint.ErrUnsupportedContentType
	}

	b := make([]byte, r.ContentLength)
	_, _ = r.Body.Read(b)
	_ = r.Body.Close()

	var dto dto
	if err := json.Unmarshal(b, &dto); err != nil {
		return nil, err
	}

	return car.New(
		dto.ID,
		dto.Make,
		dto.Model,
		dto.TrimLevel,
		dto.Color,
		dto.Year,
		dto.Category,
		dto.Mileage,
		dto.Price,
	)
}
