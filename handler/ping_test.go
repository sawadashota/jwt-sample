package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sawadashota/jwt-sample/handler"
)

func TestPingHandler(t *testing.T) {
	cases := map[string]struct {
		method       string
		expectedBody string
		expectStatus int
	}{
		"get ping": {
			method:       http.MethodGet,
			expectedBody: "Pong",
			expectStatus: http.StatusOK,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(c.method, "/ping", nil)

			handler.PingHandler(w, r)

			rw := w.Result()
			defer rw.Body.Close()

			if rw.StatusCode != c.expectStatus {
				t.Errorf("bad response status. expect: %d but actual: %d", c.expectStatus, rw.StatusCode)
			}

			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal(err)
			}

			if string(b) != c.expectedBody {
				t.Errorf("unexpected response body. expect: %s but actual: %s", c.expectedBody, string(b))
			}
		})
	}
}
