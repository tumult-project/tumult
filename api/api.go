package api

import (
	"log"
	"net/http"

	"github.com/tumult-project/tumult/api/route"
)

// Service is an API service implementation
type Service struct {
	server *http.Server
}

// Start starts the API service
func (s *Service) Start() error {
	log.Println("start API service")

	s.server = &http.Server{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", route.Home)
	mux.HandleFunc("/teapot", route.TeaPot)

	go http.ListenAndServe(":8080", mux)

	return nil
}

// Stop stops a running API service
func (s *Service) Stop() error {
	log.Println("stop API service")

	s.server.Shutdown(nil)

	return nil
}
