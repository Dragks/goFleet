package zmq

import (
	"fmt"
	"github.com/zeromq/goczmq"
	"log"
)

type Adapter struct {
	dealer *goczmq.Sock
}

func (zmqAdapter Adapter) CloseConnection() {
	zmqAdapter.dealer.Destroy()
}

func NewAdapter(connection string) (*Adapter, error) {
	dealer, err := goczmq.NewDealer(connection)
	if err != nil {
		log.Fatalf("failed to create zeromq dealer: %v", err)
	}
	return &Adapter{dealer: dealer}, nil
}

func (zmqAdapter Adapter) DoSend(value float32, sensor string) error {
	return zmqAdapter.dealer.SendFrame([]byte(fmt.Sprintf("%f", value)), goczmq.FlagNone)
}
