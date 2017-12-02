package agent

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tumult-project/tumult/api"
	"github.com/tumult-project/tumult/service"
)

// Agent ...
type Agent struct {
	APIService service.Servicer

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

	a.APIService = &api.Service{}
	a.APIService.Start()

	a.handleGracefulShutdown()
}

// Stop stops a running agent
func (a *Agent) Stop() {
	log.Println("Stopping agent...")

	a.APIService.Stop()
}

// handleGracefulShutdown enables graceful shutdown for this agent
//
// listens for two signals
// SIGTERM: generic signal used to cause program termination.
// SIGINT: signal used when the user types C-c
func (a *Agent) handleGracefulShutdown() {
	signal.Notify(a.signal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-a.signal
		a.Stop()
		a.quit <- true
	}()
	<-a.quit
}
