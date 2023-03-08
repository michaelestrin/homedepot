/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package car

import (
	"HomeDepot/internal/pkg/ui/endpoint"
	"net/http"
	"strings"
)

func idFromURLPath(path string) (string, bool) {
	p := strings.Split(strings.TrimPrefix(path, rootForGet), "/")
	if len(p) != 1 || len(p[0]) == 0 {
		return "", false
	}
	return p[0], true
}

func (o *Endpoint) get(rw http.ResponseWriter, r *http.Request) {
	id, ok := idFromURLPath(r.URL.Path)
	if !ok {
		endpoint.Respond(rw, nil, endpoint.ErrIDinURL)
		return
	}

	m, err := o.useCaseGetCar.Do(id)
	endpoint.Respond(rw, toDTO(m), err)
}
