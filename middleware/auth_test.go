package middleware_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sawadashota/jwt-sample/middleware"

	"github.com/sawadashota/jwt-sample/handler"
)

func TestAuth(t *testing.T) {
	lr := loginRequest(t)

	cases := map[string]struct {
		method       string
		token        string
		url          string
		expectedBody string
		expectStatus int
	}{
		"correct jwt token": {
			method:       http.MethodPost,
			token:        lr.Data.Token,
			url:          "/ping",
			expectedBody: "Pong",
			expectStatus: http.StatusOK,
		},
		"incorrect jwt token": {
			method:       http.MethodPost,
			token:        "this_is_invalid_token",
			url:          "/ping",
			expectedBody: "Unauthorized\n",
			expectStatus: http.StatusUnauthorized,
		},
		"empty authorization at header": {
			method:       http.MethodPost,
			token:        "",
			url:          "/ping",
			expectedBody: "Unauthorized\n",
			expectStatus: http.StatusUnauthorized,
		},
	}

	ts := httptest.NewServer(middleware.Auth(getPingHandler()))
	defer ts.Close()

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {

			var u bytes.Buffer
			u.WriteString(string(ts.URL))
			u.WriteString(c.url)

			var client http.Client
			req, err := http.NewRequest(c.method, u.String(), nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Add("Authorization", c.token)

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != c.expectStatus {
				t.Errorf("bad response status. expect: %d but actual: %d", c.expectStatus, resp.StatusCode)
			}

			if string(b) != c.expectedBody {
				t.Errorf("unexpected response body. expect: %s but actual: %s", c.expectedBody, string(b))
			}
		})
	}
}

func getPingHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	}

	return http.HandlerFunc(fn)
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
