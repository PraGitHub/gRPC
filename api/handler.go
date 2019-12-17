package api

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
)

// Server represents the gRPC server
type Server struct {
	Name string
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s from %s", in.Message, in.Sender)
	return &PingMessage{Sender: s.Name, Message: fmt.Sprintf("Hey %s, nice to here from you !", in.Sender)}, nil
}
