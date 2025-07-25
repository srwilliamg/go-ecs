package usecase

import (
	repositoryUsers "srwilliamg/app/v1/internal/repositories/users"
)

type User struct {
	userRepository repositoryUsers.UserRepository
}

func NewUser(userRepository repositoryUsers.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) GetUser(id string) (string, error) {
	user, err := u.userRepository.GetUsers()
	if err != nil {
		return "", err
	}
	return users, nil
}
