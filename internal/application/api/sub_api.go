package api

import (
	"fmt"
	"goFleet/internal/ports"
	"log"
	"net"
)

type SubApplication struct {
	zmqSub              ports.ZmqSubPort
	subscriptionHandler ports.SubscriptionHandler
}

type UDPWriter struct {
	connection net.Conn
}

func NewSubApplication(zmqSub ports.ZmqSubPort, subscriptionHandler ports.SubscriptionHandler) *SubApplication {
	return &SubApplication{subscriptionHandler: subscriptionHandler, zmqSub: zmqSub}
}

func NewUDPWriter(endpoint string) (*UDPWriter, error) {
	conn, err := net.Dial("udp", endpoint)
	if err != nil {
		return nil, err
	}
	return &UDPWriter{connection: conn}, nil
}

func (writer UDPWriter) HandleResult(address, content string) error {
	_, err := writer.connection.Write([]byte(fmt.Sprintf("%s:%s", address, content)))
	if err != nil {
		return err
	}
	return nil
}

func (writer UDPWriter) Close() {
	_ = writer.connection.Close()
}

func (app SubApplication) ReceiveAndHandle() error {
	var err error
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
