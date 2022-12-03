package api

import (
	"goFleet/internal/ports"
)

type PubApplication struct {
	zmqPubPort ports.ZmqPubPort
	sensor     Sensor
	topic      string
}

func NewPubApplication(zmqPub ports.ZmqPubPort, sensor Sensor, topic string) *PubApplication {
	return &PubApplication{zmqPubPort: zmqPub, sensor: sensor, topic: topic}
}

func (app PubApplication) ReadAndPublish() (float32, error) {
	val, err := app.sensor.ReadCurrentValue()
	if err != nil {
		return 0, err
	}

	identifier, err := app.sensor.GetIdentifier()
	if err != nil {
		return 0, err
	}

	if app.zmqPubPort != nil {
		err = app.zmqPubPort.Publish(val, identifier, app.topic)
		if err != nil {
			return 0, err
		}
	}
	return val, nil
}
