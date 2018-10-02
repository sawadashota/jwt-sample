package middleware

import (
	"net/http"
	"os"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

type Stack struct {
	mux         *http.ServeMux
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

func New(mux *http.ServeMux, mws ...middleware) *Stack {
	return &Stack{
		mux:         mux,
		middlewares: append([]middleware{logger}, mws...),
	}
}

func (s Stack) Then(h http.HandlerFunc) http.HandlerFunc {
	for i := range s.middlewares {
		h = s.middlewares[len(s.middlewares)-1-i](h)
	}
	return h
}

type RouteRegister interface {
	Pattern() string
	Handler() http.HandlerFunc
	WithMiddleware(handlerFunc http.HandlerFunc)
}

func (s Stack) Group(registers ...RouteRegister) []RouteRegister {
	for _, register := range registers {
		for i := range s.middlewares {
			register.WithMiddleware(s.middlewares[len(s.middlewares)-1-i](register.Handler()))
		}

		s.mux.HandleFunc(register.Pattern(), register.Handler())
	}

	return registers
}
