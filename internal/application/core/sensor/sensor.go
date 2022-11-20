package sensor

import "math/rand"

type Sensor struct {
	// some driver or file to read from
}

func New() *Sensor {
	return &Sensor{}
}

func (sensor Sensor) Read() (float64, error) {
	// TODO: Read actual value
	return rand.Float64() * 25, nil
}
