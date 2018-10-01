package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sawadashota/jwt-sample/handler"
)

func TestRequireTokenAuthenticationHandler(t *testing.T) {
	lr := loginRequest(t)

	cases := map[string]struct {
		method       string
		token        string
		expectStatus int
	}{
		"correct jwt token": {
			method:       http.MethodPost,
			token:        lr.Data.Token,
			expectStatus: http.StatusOK,
		},
		"incorrect jwt token": {
			method:       http.MethodPost,
			token:        "this_is_invalid_token",
			expectStatus: http.StatusUnauthorized,
		},
		"empty authorization at header": {
			method:       http.MethodPost,
			token:        "",
			expectStatus: http.StatusUnauthorized,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(c.method, "/auth", nil)
			r.Header.Add("Authorization", c.token)
			handler.RequireTokenAuthenticationHandler(w, r)

			rw := w.Result()
			defer rw.Body.Close()

			if rw.StatusCode != c.expectStatus {
				t.Errorf("bad response status. expect: %d but actual: %d", c.expectStatus, rw.StatusCode)
			}
		})
	}
}

func loginRequest(t *testing.T) *handler.LoginResponse {
	t.Helper()

	bodyMap := map[string]interface{}{
		"data": map[string]interface{}{
			"username": "test",
			"password": "test",
		},
	}

	body, err := json.Marshal(bodyMap)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))

	handler.LoginHandler(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	return readLoginResponse(t, rw.Body)
}

// readLoginResponse unmarshal LoginResponse struct from response body
func readLoginResponse(t *testing.T, r io.ReadCloser) *handler.LoginResponse {
	t.Helper()

	b, err := ioutil.ReadAll(r)

	defer r.Close()
	if err != nil {
		t.Fatal(err)
	}

	var resp handler.LoginResponse
	if err = json.Unmarshal(b, &resp); err != nil {
		t.Fatal(err)
	}

	if resp.Data.Token == "" {
		t.Fatal(fmt.Errorf("response nbody has no JWT token"))
	}

	return &resp
}
