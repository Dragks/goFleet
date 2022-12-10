package api

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"goFleet/internal/ports"
	"log"
)

type SubApplication struct {
	zmqSub              ports.ZmqSubPort
	subscriptionHandler SubscriptionHandler
}

type EndpointWriter struct {
	socket     *zmq.Socket
	identifier string
}

func NewSubApplication(zmqSub ports.ZmqSubPort, subscriptionHandler SubscriptionHandler) *SubApplication {
	return &SubApplication{subscriptionHandler: subscriptionHandler, zmqSub: zmqSub}
}

func NewEndpointWriter(endpoint string) (*EndpointWriter, error) {
	socket, err := zmq.NewSocket(zmq.STREAM)
	if err != nil {
		log.Printf("NewSocket failed %v", err)
		return nil, err
	}
	err = socket.Connect(endpoint)
	if err != nil {
		log.Printf("Connect failed %v", err)
		return nil, err
	}
	parts, err := socket.RecvMessage(0)
	if err != nil {
		log.Printf("recv failed %v", err)
		return nil, err
	}
	identifier := parts[0]

	_, err = socket.Send(identifier, zmq.SNDMORE)
	if err != nil {
		log.Printf("sending identifier failed %v", err)
		return nil, err
	}
	_, err = socket.Send("START\n", 0)
	if err != nil {
		log.Printf("START ping failed %v", err)
		return nil, err
	}

	return &EndpointWriter{socket: socket, identifier: identifier}, nil
}

func (writer EndpointWriter) HandleResult(address, content string) error {
	_, err := writer.socket.Send(writer.identifier, zmq.SNDMORE)
	if err != nil {
		log.Printf("sending identifier failed %v", err)
		return err
	}
	message := fmt.Sprintf("%s:%s\n", address, content)
	_, err = writer.socket.Send(message, 0)
	if err != nil {
		log.Printf("sending message failed %v", err)
		return err
	}

	return nil
}

func (writer EndpointWriter) Close() {
	_, _ = writer.socket.Send(writer.identifier, zmq.SNDMORE)
	_, _ = writer.socket.Send("STOP\n", 0)
	_ = writer.socket.Close()
}

func (app SubApplication) ReceiveAndHandle() error {
	address, content, err := app.zmqSub.Receive()
	if err != nil {
		return err
	}
	log.Printf("received %s on %s\n", content, address)

	err = app.subscriptionHandler.HandleResult(address, content)
	if err != nil {
		return err
	}

	return nil
}
