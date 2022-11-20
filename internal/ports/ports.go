package ports

type APIPort interface {
	GetRead() (float64, error)
}
