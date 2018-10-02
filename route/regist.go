package route

import "net/http"

type MuxHandleFunc struct {
	pattern string
	handler http.HandlerFunc
}

func regist(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return &MuxHandleFunc{
		pattern: pattern,
		handler: handler,
	}
}

func (m *MuxHandleFunc) Pattern() string {
	return m.pattern
}

func (m *MuxHandleFunc) Handler() http.HandlerFunc {
	return m.handler
}

func (m *MuxHandleFunc) WithMiddleware(h http.HandlerFunc) {
	m.handler = h
}
