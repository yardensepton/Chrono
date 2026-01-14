package services

import (
	"my-go-server/repositories"
	"my-go-server/users"
)

 type UserService struct {
    repo repositories.Repository[users.User]
}
func NewUserService(repo repositories.Repository[users.User]) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) InsertUser(userReq users.UserRequest) (users.User, error) {
	userModel := users.NewUser(userReq)
	return s.repo.Insert(userModel)
}

func (s *UserService) GetUserByID(id string) (users.User, error) {
	// Implementation for getting a user by ID
	return s.repo.GetByID(id)			
}

func (s *UserService) UpdateUser(user users.User) (users.User, error) {
	// Implementation for updating a user
	return s.repo.Update(user)
}

func DeleteUser() {
	// Implementation for deleting a user
}	