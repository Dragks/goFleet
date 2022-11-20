package api

type Sensor interface {
	Read() (float32, error)
	Identifier() (string, error)
}
