/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package car

import (
	"HomeDepot/internal/pkg/ui/endpoint"
	"net/http"
)

func (o *Endpoint) put(rw http.ResponseWriter, r *http.Request) {
	m, err := toModel(r)
	if err == nil {
		err = o.useCaseUpdateCar.Do(m)
	}
	endpoint.Respond(rw, nil, err)
}
