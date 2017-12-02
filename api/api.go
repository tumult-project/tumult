package api

import (
	"log"
	"net/http"

	"github.com/tumult-project/tumult/api/route"
)

// Service is a public API service implementation
type Service struct {
	server *http.Server
}

// Start starts the API service
func (s *Service) Start() error {
	log.Println("start API service")

	s.server = &http.Server{Addr: ":8080"}

	http.HandleFunc("/", route.Home)
	http.HandleFunc("/teapot", route.TeaPot)

	go s.server.ListenAndServe()

	return nil
}

// Stop stops a running API service
func (s *Service) Stop() error {
	log.Println("stop API service")

	s.server.Shutdown(nil)

	return nil
}
