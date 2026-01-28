package repositories

import "my-go-server/model/users"

type Repository[T any] interface {
	Insert(entity T) (T, error)
	GetByID(id string) (T, error)
	Update(entity T) (T, error)
	Delete(id string) error
}

type RefreshTokenRepository interface {
	Repository[users.RefreshToken]
	GetByToken(token string) (users.RefreshToken, error)
}
var _ RefreshTokenRepository = (*MongoRefreshTokenRepository)(nil)

type UserRepository interface {
    Repository[users.User]  
    GetByEmail(email string) (users.User, error)  
}

var _ UserRepository = (*MongoUserRepository)(nil)