package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestIndex(t *testing.T) {
	t.Run("return welcome string", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		ps := httprouter.Params{}

		Index(response, request, ps)

		got := response.Body.String()
		want := "Welcome to the Quiltspace!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
