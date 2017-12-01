package agent

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Agent ...
type Agent struct {
	APIService *http.Server

	signal chan os.Signal
	quit   chan bool
}

// NewAgent configures and creates a new agent instance
func NewAgent() (*Agent, error) {
	agent := &Agent{
		signal: make(chan os.Signal, 1),
		quit:   make(chan bool, 1),
	}

	return agent, nil
}

// Start starts an agent
func (a *Agent) Start() {
	log.Println("Starting agent...")

	signal.Notify(a.signal, syscall.SIGINT, syscall.SIGTERM)

	a.APIService = a.StartAPIService()

	go func() {
		_ = <-a.signal
		a.StopAPIService()
		a.quit <- true
	}()
	<-a.quit
}

// Stop stops a running agent
func (a *Agent) Stop() {
	log.Println("Stopping agent...")
}

// StartAPIService starts the API service for this agent
func (a *Agent) StartAPIService() *http.Server {
	log.Println("start api service")

	srv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("api service: ListenAndServe() error: %s\n", err)
		}
	}()
	return srv
}

// StopAPIService stops a running API service
func (a *Agent) StopAPIService() {
	log.Println("stop api service")
	a.APIService.Shutdown(nil)
}
