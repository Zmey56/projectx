package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("Received: %+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("Tick")
		}
	}
	fmt.Println("Server stopped")
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
