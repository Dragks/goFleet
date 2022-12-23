package zmq

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"log"
)

type Adapter struct {
	sock *zmq.Socket
}

type PubAdapter struct {
	Adapter
}

type SubAdapter struct {
	Adapter
	topic string
}

func (adapter Adapter) Close() {
	_ = adapter.sock.Close()
}

func NewPubAdapter(endpoint string) (*PubAdapter, error) {
	log.Printf("Build socket to publish")
	publisher, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		return nil, err
	}

	err = publisher.Bind(endpoint) // i.e. "tcp://*:5563"
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully built publisher socket: %s", endpoint)
	return &PubAdapter{Adapter: Adapter{sock: publisher}}, nil
}

func NewSubAdapter(connection, topic string) (*SubAdapter, error) {
	log.Printf("Build socket for subscription")
	subscriber, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return nil, err
	}
	log.Printf("Connecting socket for subscription")
	err = subscriber.Connect(connection) // i.e. "tcp://localhost:5563"
	if err != nil {
		return nil, err
	}
	err = subscriber.SetSubscribe(topic)
	if err != nil {
		return nil, err
	}

	log.Printf("Sucessfully built subscriber socket '%s' for topic '%s'", connection, topic)
	return &SubAdapter{topic: topic, Adapter: Adapter{sock: subscriber}}, nil
}

func (pubAdapter PubAdapter) Publish(value float32, sensorId, topic string) error {
	log.Printf("Publish next message to all subscribers")
	_, err := pubAdapter.sock.Send(topic, zmq.SNDMORE)
	if err != nil {
		return err
	}
	_, err = pubAdapter.sock.Send(fmt.Sprintf("%f", value), 0)
	if err != nil {
		return err
	}
	log.Printf("publisher sent '%f' on topic '%s' with sensor '%s'", value, topic, sensorId)
	return err
}

func (subAdapter SubAdapter) Receive() (string, string, error) {
	log.Printf("Receive next message from subscription")
	//  Read envelope with address
	address, err := subAdapter.sock.Recv(0)
	if err != nil {
		return "", "", err
	}
	log.Printf("Received address '%s' from subscription", address)
	//  Read message contents
	content, err := subAdapter.sock.Recv(0)
	if err != nil {
		return "", "", err
	}
	log.Printf("subscriber received '%s' from '%s' on topic '%s'", content, address, subAdapter.topic)
	return address, content, nil
}
