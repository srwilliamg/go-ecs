package usecase

import (
	"srwilliamg/app/v1/internal/application/dto"
	repositoryUsers "srwilliamg/app/v1/internal/repositories/users"
)

type User struct {
	userRepository repositoryUsers.UserRepository
}

func NewUser(userRepository *repositoryUsers.UserRepository) *User {
	return &User{
		userRepository: *userRepository,
	}
}

func (u *User) GetUser() ([]dto.User, error) {
	users, err := u.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	var userDTOs []dto.User

	if users != nil {
		for _, user := range users.Rows {
			nu := dto.FromEntity(&user)
			userDTOs = append(userDTOs, *nu)
		}

	}
	return userDTOs, nil
}
