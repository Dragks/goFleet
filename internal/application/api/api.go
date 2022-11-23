package api

import (
	"goFleet/internal/ports"
)

type Application struct {
	db     ports.DbPort
	zmq    ports.ZmqPort
	sensor Sensor
}

func NewApplication(zmq ports.ZmqPort, db ports.DbPort, sensor Sensor) *Application {
	return &Application{zmq: zmq, db: db, sensor: sensor}
}

func NewDbApplication(db ports.DbPort, sensor Sensor) *Application {
	return &Application{db: db, sensor: sensor}
}

func NewZmqApplication(zmq ports.ZmqPort, sensor Sensor) *Application {
	return &Application{zmq: zmq, sensor: sensor}
}

func (app Application) GetRead() (float32, error) {
	val, err := app.sensor.Read()
	if err != nil {
		return 0, err
	}

	identifier, err := app.sensor.Identifier()
	if err != nil {
		return 0, err
	}

	if app.db != nil {
		err = app.db.LogHistory(val, identifier)
		if err != nil {
			return 0, err
		}
	}

	if app.zmq != nil {
		err = app.zmq.DoSend(val, identifier)
		if err != nil {
			return 0, err
		}
	}
	return val, nil
}
