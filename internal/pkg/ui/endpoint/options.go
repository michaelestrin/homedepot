/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package endpoint

import (
	"net/http"
	"strings"
)

var (
	HeaderAccessControlAllowOriginKey   = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowOriginValue = "*"

	HeaderAccessControlAllowMethodsKey = "Access-Control-Allow-Methods"

	HeaderAccessControlAllowHeadersKey   = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowHeadersValue = HeaderContentTypeKey
)

func OptionsHeaders(verbs []string) map[string]string {
	return map[string]string{
		HeaderAccessControlAllowOriginKey:  HeaderAccessControlAllowOriginValue,
		HeaderAccessControlAllowMethodsKey: strings.Join(append(verbs, http.MethodOptions), ","),
		HeaderAccessControlAllowHeadersKey: HeaderAccessControlAllowHeadersValue,
	}
}

func Options(rw http.ResponseWriter, verbs []string) {
	for k, v := range OptionsHeaders(verbs) {
		rw.Header().Set(k, v)
	}
}
