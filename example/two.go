package main

import (
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	log.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.REQ)
	err := s.Connect("tcp://zmq:5555")
	if err != nil {
		return
	}

	// Do 10 requests, waiting each time for a response
	for i := 0; i < 10; i++ {
		log.Printf("Sending request %d...\n", i)
		_, err := s.Send("Hello", 0)
		if err != nil {
			return
		}

		msg, _ := s.Recv(0)
		log.Printf("Received reply %d [ %s ]\n", i, msg)
	}
}
