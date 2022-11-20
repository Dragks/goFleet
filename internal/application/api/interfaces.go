package api

type Sensor interface {
	Read() (float64, error)
}
