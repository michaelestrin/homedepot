/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package metrics

import (
	"fmt"
	"net/http"
)

func Requests(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("> %5s %s\n", r.Method, r.RequestURI)
		h(rw, r)
		fmt.Printf("< %5s %s\n", r.Method, r.RequestURI)
	}
}
