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
	log.Println("Creating NewSocket")
	socket, err := zmq.NewSocket(zmq.STREAM)
	if err != nil {
		log.Printf("NewSocket failed %v", err)
		return nil, err
	}
	log.Printf("Connecting to the Endpoint: %s", endpoint)
	err = socket.Connect(endpoint)
	if err != nil {
		log.Printf("Connect failed %v", err)
		return nil, err
	}
	log.Printf("Receive first message from the endpoint")
	parts, err := socket.RecvMessage(0)
	if err != nil {
		log.Printf("recv failed %v", err)
		return nil, err
	}
	identifier := parts[0]
	log.Printf("Successfully received the identifier: %s", identifier)

	log.Printf("Send first message to the endpoint")
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

	log.Printf("Endpoint built successfully")
	return &EndpointWriter{socket: socket, identifier: identifier}, nil
}

func (writer EndpointWriter) HandleResult(address, content string) error {
	log.Printf("Handle result with address '%s' with content '%s'", address, content)
	_, err := writer.socket.Send(writer.identifier, zmq.SNDMORE)
	if err != nil {
		log.Printf("sending identifier failed %v", err)
		return err
	}
	log.Printf("Sending message to the endpoint")
	message := fmt.Sprintf("%s:%s\n", address, content)
	_, err = writer.socket.Send(message, 0)
	if err != nil {
		log.Printf("sending message failed %v", err)
		return err
	}
	log.Printf("Message sent successfully")

	return nil
}

func (writer EndpointWriter) Close() {
	_, _ = writer.socket.Send(writer.identifier, zmq.SNDMORE)
	_, _ = writer.socket.Send("STOP\n", 0)
	_ = writer.socket.Close()
}

func (app SubApplication) ReceiveAndHandle() error {
	log.Printf("Receive and handle next content")
	address, content, err := app.zmqSub.Receive()
	if err != nil {
		return err
	}
	log.Printf("Received %s on %s\n", content, address)

	err = app.subscriptionHandler.HandleResult(address, content)
	if err != nil {
		return err
	}

	return nil
}
