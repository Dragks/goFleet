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
	publisher, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		return nil, err
	}

	err = publisher.Bind(endpoint) // i.e. "tcp://*:5563"
	if err != nil {
		return nil, err
	}

	return &PubAdapter{Adapter: Adapter{sock: publisher}}, nil
}

func NewSubAdapter(connection, topic string) (*SubAdapter, error) {
	subscriber, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return nil, err
	}
	err = subscriber.Connect(connection) // i.e. "tcp://localhost:5563"
	if err != nil {
		return nil, err
	}
	err = subscriber.SetSubscribe(topic)
	if err != nil {
		return nil, err
	}

	return &SubAdapter{topic: topic, Adapter: Adapter{sock: subscriber}}, nil
}

func (pubAdapter PubAdapter) Publish(value float32, sensorId, topic string) error {
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
	//  Read envelope with address
	address, err := subAdapter.sock.Recv(0)
	if err != nil {
		return "", "", err
	}
	//  Read message contents
	content, err := subAdapter.sock.Recv(0)
	if err != nil {
		return "", "", err
	}
	log.Printf("subscriber received '%s' from '%s' on topic '%s'", content, address, subAdapter.topic)
	return address, content, nil
}
