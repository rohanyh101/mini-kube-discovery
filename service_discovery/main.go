package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/roh4nyh/service_discovery/cache"
)

const port = 5555

func main() {

	cache.Init()

	server := rpc.NewServer()

	// register rpc methods...
	if err := server.RegisterName("Discovery", new(Discovery)); err != nil {
		log.Fatal(err)
	}

	// tcp listener
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	// defer listener.Close()

	server.Accept(listener)
}
