package quilts

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestIndex(t *testing.T) {
	t.Run("return Index of all quilts", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/quilts", nil)
		response := httptest.NewRecorder()
		params := httprouter.Params{}

		Index(response, request, params)

		got := response.Code
		want := 200

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestAbout(t *testing.T) {
	t.Run("return about page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/about", nil)
		response := httptest.NewRecorder()
		params := httprouter.Params{}

		About(response, request, params)

		got := response.Code
		want := 200

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestHome(t *testing.T) {
	t.Run("return home page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		params := httprouter.Params{}

		Home(response, request, params)

		got := response.Code
		want := 200

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
