package server

import (
	"context"

	protos "../protos/greet"
	"github.com/hashicorp/go-hclog"
)

// Greeting is a gRPC server it implements the methods defined by the GreetingServer interface
type Greeting struct {
	log hclog.Logger
}

// NewGreeting creates a new Greeting server
func NewGreeting(l hclog.Logger) *Greeting {
	return &Greeting{l}
}

// GetPhrase implements the GreetingServer
func (c *Greeting) GetPhrase(ctx context.Context, rr *protos.PhraseRequest) (*protos.PhraseResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetName())
	return &protos.PhraseResponse{Greet: "Hey there How are you"}, nil
}
