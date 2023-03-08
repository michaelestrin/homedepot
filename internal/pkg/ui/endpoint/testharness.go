/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package endpoint

import (
	"HomeDepot/internal/pkg/ui/contract"
	"io"
	"net/http"
	"net/http/httptest"
)

type testHarness struct {
	endpoint contract.Endpoint
	mux      *http.ServeMux
}

func newTestHarness(endpoint contract.Endpoint) *testHarness {
	return &testHarness{
		endpoint: endpoint,
		mux:      http.NewServeMux(),
	}
}

func (o *testHarness) Attach(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	o.mux.Handle(pattern, http.HandlerFunc(handler))
}

func (o *testHarness) Serve(method, url string, headers map[string]string, body io.Reader) *httptest.ResponseRecorder {
	o.endpoint.Register(o)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, body)
	if headers != nil {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
	}
	o.mux.ServeHTTP(w, r)
	return w
}
