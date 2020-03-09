package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	
	"github.com/ridwanakf/nadc-intro-to-rpc/server/internal"
)

func main() {
	var api = new(internal.API)

	err := rpc.Register(api)
	if err != nil {
		log.Fatal("[Error while registering API] ", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("[Listener error] ", err)
	}

	log.Printf("Serving rpc on port :%d", 4040)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("[Error serving] ", err)
	}
}
