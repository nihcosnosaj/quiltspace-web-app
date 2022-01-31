package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestQuiltsIndex(t *testing.T) {
	t.Run("return welcome string", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		ps := httprouter.Params{}

		quiltsIndex(response, request, ps)

		got := response.Body.String()
		want := "Welcome to the Quiltspace!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
