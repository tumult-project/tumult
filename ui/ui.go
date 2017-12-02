package ui

import (
	"log"
	"net/http"

	"github.com/tumult-project/tumult/ui/route"
)

// Service is a UI service implementation
type Service struct {
	server *http.Server
}

// Start starts the API service
func (s *Service) Start() error {
	log.Println("start UI service")

	s.server = &http.Server{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", route.Home)

	go http.ListenAndServe(":8081", mux)

	return nil
}

// Stop stops a running API service
func (s *Service) Stop() error {
	log.Println("stop UI service")

	s.server.Shutdown(nil)

	return nil
}
