package route

import (
	"io"
	"net/http"
)

// Home is the home route for the API service
func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "API service home\n")
}
