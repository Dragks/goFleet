package api

type Sensor interface {
	ReadCurrentValue() (float32, error)
	GetIdentifier() (string, error)
}
type SubscriptionHandler interface {
	Close()
	HandleResult(string, string) error
}
