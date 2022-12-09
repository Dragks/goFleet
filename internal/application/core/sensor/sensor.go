package sensor

import (
	"fmt"
	"math/rand"
	"time"
)

type Sensor struct {
	identifier string
	// some driver or file to read from
}

type Mock struct {
	Sensor
}

func New(identifier string) (*Sensor, error) {
	return nil, fmt.Errorf("failed to create sensor %s: not implemented", identifier)
}

func NewMock() (*Mock, error) {
	return &Mock{Sensor{identifier: "mock"}}, nil
}

func (sensor Mock) ReadCurrentValue() (float32, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() * 25, nil
}

func (sensor Sensor) GetIdentifier() (string, error) {
	return sensor.identifier, nil
}

func (sensor Sensor) ReadCurrentValue() (float32, error) {
	// TODO: ReadCurrentValue actual value
	return 0, nil
}
