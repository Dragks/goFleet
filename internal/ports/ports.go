package ports

type APIPort interface {
	GetRead() (float32, error)
}

type DbPort interface {
	CloseConnection()
	LogHistory(value float32, sensor string) error
}
