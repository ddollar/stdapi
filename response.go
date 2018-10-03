package stdapi

import "net/http"

type Response struct {
	http.ResponseWriter
	Code int
}

func (r Response) Flush() {
	if f, ok := r.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (r Response) WriteHeader(code int) {
	r.Code = code
	r.ResponseWriter.WriteHeader(code)
}
