package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestQuiltsIndex(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/quilts", nil)
	ps := httprouter.Params{}

	quiltsIndex(wr, req, ps)

	if wr.Code != http.StatusOK {
		t.Errorf("got http status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Quilt Index:") {
		t.Errorf("Response body doesn't contain 'Quilt Index:")
	}
}

func TestQuiltsCreateForm(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/quilts/create", nil)
	ps := httprouter.Params{}

	quiltsCreateForm(wr, req, ps)

	if wr.Code != http.StatusOK {
		t.Errorf("got http status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), "Create New Quilt") {
		t.Errorf("response body doesn't contain 'Create New Quilt'")
	}
}
