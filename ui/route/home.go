package route

import (
	"io"
	"net/http"
)

// Home is the home route for the UI service
func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "UI service home\n")
}
