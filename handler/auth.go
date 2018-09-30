package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// RequireTokenAuthenticationHandler verify token
func RequireTokenAuthenticationHandler(w http.ResponseWriter, r *http.Request) {

	verifyBytes, err := ioutil.ReadFile(idRsaPublicPath)
	if err != nil {
		panic(err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodRSA)
		if !err {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else {
			return verifyKey, nil
		}
	})

	if err == nil && token.Valid {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
