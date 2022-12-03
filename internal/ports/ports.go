package ports

type PubAPIPort interface {
	ReadAndPublish() (float32, error)
}

type SubAPIPort interface {
	SubscribeAndSave() (float32, error)
}

type DbPort interface {
	Close()
	LogHistory(value float32, address string) error
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
