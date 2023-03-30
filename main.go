package main

import (
	"github.com/claudiumocanu/go-mtls-example/alpha"
	"github.com/claudiumocanu/go-mtls-example/bravo"
	"github.com/claudiumocanu/go-mtls-example/charlie"
)

func main() {

	// Start all servers
	go alpha.StartServer()
	go bravo.StartServer()
	go charlie.StartServer()

	// never ending...
	c := make(chan struct{})
	<-c
}
