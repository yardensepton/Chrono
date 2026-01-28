package repositories

type Repository[T any] interface {
	Insert(entity T) (T, error)
	GetByID(id string) (T, error)
	Update(entity T) (T, error)
	Delete(id string) error
}
