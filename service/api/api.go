package api

import (
	"log"
	"net/http"

	"github.com/tumult-project/tumult/service/api/route"
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

	go s.server.ListenAndServe()

	return nil
}

// Stop stops a running API service
func (s *Service) Stop() error {
	log.Println("stop API service")

	s.server.Shutdown(nil)

	return nil
}
