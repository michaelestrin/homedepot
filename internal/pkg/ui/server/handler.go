/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package server

import (
	"HomeDepot/internal/pkg/ui/contract"
	"net/http"
)

type Server struct {
	addr       string
	mux        *http.ServeMux
	decorators []func(h http.HandlerFunc) http.HandlerFunc
}

func New(addr string, decorators ...func(h http.HandlerFunc) http.HandlerFunc) *Server {
	return &Server{
		addr:       addr,
		mux:        http.NewServeMux(),
		decorators: decorators,
	}
}

func (o *Server) Attach(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	p := handler
	for _, decorator := range o.decorators {
		p = decorator(p)
	}
	o.mux.HandleFunc(pattern, p)
}

func (o *Server) Run(endpoints []contract.Endpoint) {
	s := http.Server{
		Addr:    o.addr,
		Handler: o.mux,
	}

	for _, endpoint := range endpoints {
		endpoint.Register(o)
	}

	_ = s.ListenAndServe()
	defer func() {
		_ = s.Close()
	}()
}
