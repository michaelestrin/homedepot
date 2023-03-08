/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package cars

import (
	"HomeDepot/internal/pkg/ui/endpoint"
	"net/http"
)

func (o *Endpoint) get(rw http.ResponseWriter, _ *http.Request) {
	m, err := o.useCaseGetCars.Do()
	endpoint.Respond(rw, toDTO(m), err)
}
