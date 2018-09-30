package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ResponseBase struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type LoginResponse struct {
	ResponseBase
	Data LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
	Token     string `json:"token"`
}

type LoginRequest struct {
	Data LoginRequestData `json:"data"`
}

type LoginRequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler issue JWT token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req, err := readLoginRequest(r.Body)

	signBytes, err := ioutil.ReadFile(idRsaPath)
	if err != nil {
		loginErrorResponse(w, err)
		return
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		loginErrorResponse(w, err)
		return
	}

	// TODO: read user data from Database
	// TODO: make password be encrypted
	if req.Data.Username != "test" || req.Data.Password != "test" {
		loginInvalidResponse(w)
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		loginErrorResponse(w, err)
		return
	}

	resp := &LoginResponse{
		ResponseBase: ResponseBase{
			Status: http.StatusOK,
		},
		Data: LoginResponseData{
			Name:      "TEST",
			Token:     tokenString,
			Timestamp: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	b, err := json.Marshal(resp)
	if err != nil {
		loginErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// readLoginRequest unmarshal LoginRequest struct from request body
func readLoginRequest(r io.ReadCloser) (*LoginRequest, error) {
	b, err := ioutil.ReadAll(r)

	defer r.Close()
	if err != nil {
		return nil, err
	}

	var req LoginRequest
	if err = json.Unmarshal(b, &req); err != nil {
		return nil, err
	}

	return &req, nil
}

func loginErrorResponse(w http.ResponseWriter, err error) {
	resp := &LoginResponse{
		ResponseBase: ResponseBase{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		},
	}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}

func loginInvalidResponse(w http.ResponseWriter) {
	resp := &LoginResponse{
		ResponseBase: ResponseBase{
			Status: http.StatusForbidden,
			Error:  "username or password is incorrect",
		},
	}

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	w.Write(b)
}
