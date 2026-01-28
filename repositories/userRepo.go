package repositories

import (
	"my-go-server/users"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

var _ Repository[users.User] = (*PostgresUserRepository)(nil)

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Insert(user users.User) (users.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *PostgresUserRepository) GetByID(id string) (users.User, error) {
	var user users.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *PostgresUserRepository) Update(user users.User) (users.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *PostgresUserRepository) Delete(id string) error {
	return r.db.Delete(&users.User{}, "id = ?", id).Error
}
