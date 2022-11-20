package api

import (
	"goFleet/internal/ports"
)

type Application struct {
	db     ports.DbPort
	sensor Sensor
}

func NewApplication(db ports.DbPort, sensor Sensor) *Application {
	return &Application{db: db, sensor: sensor}
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
	err = app.db.LogHistory(val, identifier)
	if err != nil {
		return 0, err
	}
	return val, nil
}
