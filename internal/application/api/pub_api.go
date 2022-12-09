package api

import (
	"goFleet/internal/ports"
	"log"
)

type PubApplication struct {
	zmqPubPort ports.ZmqPubPort
	sensor     Sensor
	topic      string
}

func NewPubApplication(zmqPub ports.ZmqPubPort, sensor Sensor, topic string) *PubApplication {
	return &PubApplication{zmqPubPort: zmqPub, sensor: sensor, topic: topic}
}

func (app PubApplication) ReadAndPublish() error {
	val, err := app.sensor.ReadCurrentValue()
	if err != nil {
		return err
	}
	log.Printf("got sensor read with value: %f\n", val)

	identifier, err := app.sensor.GetIdentifier()
	if err != nil {
		return err
	}

	if app.zmqPubPort != nil {
		err = app.zmqPubPort.Publish(val, identifier, app.topic)
		if err != nil {
			return err
		}
	}
	return nil
}
