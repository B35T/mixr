package mixr

import "net/http"

type NewMixr interface {
	Get(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler
	Post(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler
	Put(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler
	Patch(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler
	Del(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler
}

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

func Get(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == MethodGet {
			endpoint(w, r)
		} else {
			w.WriteHeader(500)
		}
	})
}

func Post(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == MethodPost {
			endpoint(w, r)
		} else {
			w.WriteHeader(500)
		}
	})
}

func Put(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == MethodPut {
			endpoint(w, r)
		} else {
			w.WriteHeader(500)
		}
	})
}

func Patch(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == MethodPatch {
			endpoint(w, r)
		} else {
			w.WriteHeader(500)
		}
	})
}

func Del(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == MethodDelete {
			endpoint(w, r)
		} else {
			w.WriteHeader(500)
		}
	})
}
