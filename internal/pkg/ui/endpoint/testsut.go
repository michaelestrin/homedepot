/**
 * Copyright(c) 2023 Michael Estrin.  All rights reserved.
 */

package endpoint

import (
	appContract "HomeDepot/internal/pkg/application/car/contract"
	"HomeDepot/internal/pkg/domain/car"
	"HomeDepot/internal/pkg/infrastructure/car/repository"
	uiContract "HomeDepot/internal/pkg/ui/contract"
	v1 "HomeDepot/internal/pkg/ui/endpoint/v1"
	"bytes"
	"io"
	"testing"
)

var (
	NoRepositoryInit      []*car.Model      = nil
	NoRepositoryAssertion []*car.Model      = nil
	NoHeaderInit          map[string]string = nil
	AllowOriginHeader                       = map[string]string{HeaderAccessControlAllowOriginKey: HeaderAccessControlAllowOriginValue}
	NoBodyInit            io.Reader         = nil
	EmptyBody                               = ""

	HeaderContentTypeJSON    = map[string]string{HeaderContentTypeKey: HeaderContentTypeValueJSON}
	HeaderContentTypeInvalid = map[string]string{HeaderContentTypeKey: "invalid"}
)

type Test struct {
	Name            string
	RepoPre         []*car.Model
	Method          string
	URL             string
	Headers         map[string]string
	Body            io.Reader
	ExpectedStatus  int
	ExpectedHeaders map[string]string
	ExpectedBody    string
	RepoPost        []*car.Model
}

type GetEndpoint func(r appContract.Repository) uiContract.Endpoint

type SUT struct {
	getEndpoint GetEndpoint
	tests       []Test
}

func NewSUT(getEndpoint GetEndpoint, tests []Test) *SUT {
	return &SUT{
		getEndpoint: getEndpoint,
		tests:       tests,
	}
}

func (o SUT) Run(t *testing.T) {
	for _, test := range o.tests {
		r := repository.NewCar()
		if test.RepoPre != nil {
			for _, m := range test.RepoPre {
				_ = r.Create(m)
			}
		}
		sut := newTestHarness(o.getEndpoint(r))

		w := sut.Serve(test.Method, test.URL, test.Headers, test.Body)

		if w.Code != test.ExpectedStatus {
			t.Errorf("%s %s - status expected: %v, actual: %v", test.URL, test.Name, test.ExpectedStatus, w.Code)
		}

		if test.ExpectedHeaders != nil {
			for headerKey, headerValue := range test.ExpectedHeaders {
				actualValue := w.Header().Get(headerKey)
				switch {
				case actualValue == "":
					t.Errorf("%s %s - Header %v missing", test.URL, test.Name, headerKey)
				case actualValue != headerValue:
					t.Errorf("%s %s - Header expected: %v: %v, actual: %v: %v",
						test.URL,
						test.Name,
						headerKey,
						headerValue,
						headerKey,
						actualValue,
					)
				}
			}
		}

		b := w.Body.Bytes()
		if bytes.Compare(b, []byte(test.ExpectedBody)) != 0 {
			t.Errorf("%s %s - Body expected: %v, actual: %v", test.URL, test.Name, test.ExpectedBody, string(b))
		}

		if test.RepoPost != nil {
			models, _ := r.GetAll()
			if len(models) != len(test.RepoPost) {
				t.Errorf("%s %s - repo expected: %v, actual: %v", test.URL, test.Name, len(test.RepoPost), len(models))
			}
			for _, expected := range test.RepoPost {
				actual, err := r.GetByID(expected.ID)
				if err != nil {
					t.Errorf("%s %s - repo missing expected: %v", test.URL, test.Name, expected)
					continue
				}

				expectedJSON := v1.AsJSON(expected)
				actualJSON := v1.AsJSON(actual)
				if actualJSON != expectedJSON {
					t.Errorf("%s %s - repo expected: %v, actual: %v", test.URL, test.Name, expectedJSON, actualJSON)
				}
			}
		}
	}
}
