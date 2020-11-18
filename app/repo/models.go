package repo

// Repository is exported type
type Repository interface {
	GetAll() (string, error)
	Ping() (int, error)
}
