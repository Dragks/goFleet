package api

type Sensor interface {
	ReadCurrentValue() (float32, error)
	GetIdentifier() (string, error)
}
