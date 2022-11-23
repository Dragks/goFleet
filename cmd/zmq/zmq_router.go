package main

import (
	"fmt"
	"github.com/zeromq/goczmq"
	"log"
	"os"
	"strconv"
)

func main() {
	var err error
	fmt.Println("application is starting...")

	port, err := strconv.Atoi(os.Getenv("ZMQ_PORT"))
	if err != nil {
		port = 5555
	}

	// Create a router socket and bind it to a port.
	router, err := goczmq.NewRouter(fmt.Sprintf("tcp://*:%d", port))
	if err != nil {
		log.Fatalf("failed to create zeromq router: %v", err)
	}
	defer router.Destroy()

	log.Println("router created and bound")

	// Receive the message. Here we call RecvMessage, which
	// will return the message as a slice of frames ([][]byte).
	// Since this is a router socket that support async
	// request / reply, the first frame of the message will
	// be the routing frame.
	request, err := router.RecvMessage()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("router received '%s' from '%v'", request[1], request[0])

	fmt.Println("application is closing...")
}
