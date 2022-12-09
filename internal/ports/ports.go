package ports

type PubAPIPort interface {
	ReadAndPublish() (float32, error)
}

type SubAPIPort interface {
	ReceiveAndHandle() error
}

type ZmqPort interface {
	Close()
}

type ZmqPubPort interface {
	ZmqPort
	Publish(value float32, sensorId, topic string) error
}
type ZmqSubPort interface {
	ZmqPort
	Receive() (string, string, error)
}
type SubscriptionHandler interface {
	Close()
	HandleResult(string, string) error
}
