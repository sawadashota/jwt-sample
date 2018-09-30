package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sawadashota/jwt-sample/handler"
)

func TestLoginHandler(t *testing.T) {
	cases := map[string]struct {
		method       string
		bodyMap      map[string]interface{}
		expectStatus int
	}{
		"correct username password": {
			method: http.MethodPost,
			bodyMap: map[string]interface{}{
				"data": map[string]interface{}{
					"username": "test",
					"password": "test",
				},
			},
			expectStatus: http.StatusOK,
		},
		"incorrect username password": {
			method: http.MethodPost,
			bodyMap: map[string]interface{}{
				"data": map[string]interface{}{
					"username": "hoge",
					"password": "fuga",
				},
			},
			expectStatus: http.StatusForbidden,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			body, err := json.Marshal(c.bodyMap)
			if err != nil {
				t.Fatal(err)
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(c.method, "/login", bytes.NewReader(body))

			handler.LoginHandler(w, r)

			rw := w.Result()
			defer rw.Body.Close()

			if rw.StatusCode != c.expectStatus {
				t.Errorf("bad response status. expect: %d but actual: %d", c.expectStatus, rw.StatusCode)
			}

			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal(err)
			}

			var lr handler.LoginResponse
			buf := bytes.NewReader([]byte(b))
			dec := json.NewDecoder(buf)

			if err := dec.Decode(&lr); err != nil {
				t.Fatal(err)
			}
		})
	}
}
