package route

import "net/http"

func Post(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodPost, handler))
}

func Get(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodGet, handler))
}

func Options(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodOptions, handler))
}

func Head(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodHead, handler))
}

func Connect(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodConnect, handler))
}

func Delete(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodDelete, handler))
}

func Patch(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodPatch, handler))
}

func Put(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodPut, handler))
}

func Trace(pattern string, handler http.HandlerFunc) *MuxHandleFunc {
	return regist(pattern, isSameMethod(http.MethodTrace, handler))
}

func isSameMethod(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Bad Request Method", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
