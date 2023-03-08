/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package endpoint

import (
	"HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/domain"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrMethodNotAllowed       = errors.New("method not allowed")
	ErrContentLengthIsZero    = errors.New("content length is zero")
	ErrUnsupportedContentType = errors.New("missing or unsupported content type")
	ErrIDinURL                = errors.New("cannot parse id from url")

	HeaderContentTypeKey       = "Content-Type"
	HeaderContentTypeValueJSON = "application/json"
)

func errToStatus(err error) int {
	switch {
	case errors.Is(err, ErrMethodNotAllowed):
		return http.StatusMethodNotAllowed
	case errors.Is(err, ErrIDinURL) ||
		errors.Is(err, ErrContentLengthIsZero) ||
		errors.Is(err, contract.ErrResourceAlreadyExists) ||
		errors.Is(err, domain.ErrDomainConsistency):
		return http.StatusBadRequest
	case errors.Is(err, contract.ErrResourceDoesNotExist):
		return http.StatusNotFound
	case errors.Is(err, ErrUnsupportedContentType):
		return http.StatusUnsupportedMediaType
	default:
		return http.StatusInternalServerError
	}
}

func Respond(rw http.ResponseWriter, body any, err error) {
	rw.Header().Set(HeaderAccessControlAllowOriginKey, HeaderAccessControlAllowOriginValue)
	switch {
	case err == nil && body != nil:
		rw.Header().Set(HeaderContentTypeKey, HeaderContentTypeValueJSON)

		var b []byte
		b, err = json.Marshal(body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = rw.Write(b)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}

	case err != nil:
		rw.Header().Set(HeaderContentTypeKey, HeaderContentTypeValueJSON)
		rw.WriteHeader(errToStatus(err))

		b, _ := json.Marshal(
			struct {
				Error string `json:"error"`
			}{
				err.Error(),
			},
		)
		_, _ = rw.Write(b)
	}
}
