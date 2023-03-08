/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package contract

import "net/http"

type Server interface {
	Attach(pattern string, handler func(http.ResponseWriter, *http.Request))
}
