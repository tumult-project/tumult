package service

// Servicer is an interface that defines an agent service
type Servicer interface {
	Start() error
	Stop() error
}
