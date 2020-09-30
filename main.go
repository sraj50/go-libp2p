package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {
	// create a background context that never cancels
	ctx := context.Background()

	// start libp2p node with settings configured to listen on
	// TCP port 2000 on loopback interface
	node, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"))
	if err != nil {
		panic(err)
	}

	// print node listening address
	fmt.Println("Listen addresses: ", node.Addrs())

	// shut down ode
	if err := node.Close; err != nil {
		panic(err)
	}
}