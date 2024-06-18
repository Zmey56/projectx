package main

import (
	"guthub.com/zmey56/pojectx/network"
	"time"
)

// Server
// Transport => tcp, udp,
// Block
// Tx

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage("LOCAL", []byte("Hello world"))
			time.Sleep(1 * time.Second)
		}

	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
