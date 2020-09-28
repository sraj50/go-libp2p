package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {
	// create a background context that never cancels
	ctx := context.Background()

	// start libp2p node with default settings
	node, err := libp2p.New(ctx)
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