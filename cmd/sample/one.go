package main

import (
	"log"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)
	s.Bind("tcp://*:5555")

	i := 0
	for {
		i++
		// Wait for next request from client
		msg, _ := s.Recv(0)
		log.Printf("[%d] Received %s\n", i, msg)

		// Do some 'work'
		time.Sleep(time.Second * 1)

		// Send reply back to client
		s.Send("World "+strconv.Itoa(i), 0)
	}
}
