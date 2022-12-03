package api

import (
	"goFleet/internal/ports"
	"strconv"
)

type SubApplication struct {
	db     ports.DbPort
	zmqSub ports.ZmqSubPort
}

func NewSubApplication(zmqSub ports.ZmqSubPort, db ports.DbPort) *SubApplication {
	return &SubApplication{db: db, zmqSub: zmqSub}
}

func (app SubApplication) SubscribeAndSave() (float32, error) {
	var err error
	address, content, err := app.zmqSub.Receive()
	if err != nil {
		return 0, err
	}

	parse, err := strconv.ParseFloat(content, 32)
	value := float32(parse)

	err = app.db.LogHistory(value, address)
	if err != nil {
		return value, err
	}

	return value, nil
}
