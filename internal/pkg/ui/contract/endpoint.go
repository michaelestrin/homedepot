/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package contract

type Endpoint interface {
	Register(s Server)
}
