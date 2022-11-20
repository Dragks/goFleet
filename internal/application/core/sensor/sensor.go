package sensor

import "math/rand"

type Sensor struct {
	identifier string
	// some driver or file to read from
}

func New(identifier string) *Sensor {
	return &Sensor{identifier: identifier}
}

func (sensor Sensor) Read() (float32, error) {
	// TODO: Read actual value
	return rand.Float32() * 25, nil
}

func (sensor Sensor) Identifier() (string, error) {
	return sensor.identifier, nil
}
