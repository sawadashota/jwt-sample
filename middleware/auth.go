package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		verifyBytes, err := ioutil.ReadFile(idRsaPublicPath)
		if err != nil {
			http.Error(w, "InternalServerError", http.StatusInternalServerError)
			return
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			http.Error(w, "InternalServerError", http.StatusInternalServerError)
			return
		}

		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			_, err := token.Method.(*jwt.SigningMethodRSA)
			if !err {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			} else {
				return verifyKey, nil
			}
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
