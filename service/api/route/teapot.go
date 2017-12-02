package route

import (
	"io"
	"net/http"
)

// Teapot is a route for the API service returns a HTTP 418 status code
func Teapot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	io.WriteString(w, "I'm a teapot\n")
}
