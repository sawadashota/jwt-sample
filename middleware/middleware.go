package middleware

import (
	"net/http"
	"os"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

type Stack struct {
	middlewares []middleware
}

var (
	idRsaPublicPath string
)

const (
	IdRsaPathPublicDefault = "./certs/id_rsa.pub.pkcs8"
)

func init() {
	idRsaPublicPath = getEnvString("ID_RSA_PUBLIC_PATH", IdRsaPathPublicDefault)
}
func getEnvString(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		v = defaultValue
	}

	return v
}

func New(mws ...middleware) Stack {
	return Stack{append([]middleware(nil), mws...)}
}

func (m Stack) Then(h http.HandlerFunc) http.HandlerFunc {
	for i := range m.middlewares {
		h = m.middlewares[len(m.middlewares)-1-i](h)
	}
	return h
}
