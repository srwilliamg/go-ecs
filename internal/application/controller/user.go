package controller

import (
	"net/http"
	"srwilliamg/app/v1/internal/application/dto"
	"srwilliamg/app/v1/internal/application/request"
	"srwilliamg/app/v1/internal/interfaces/logger"
)

type UserUseCase interface {
	GetUser() ([]dto.User, error)
}

type UserController struct {
	userUseCase UserUseCase
}

func NewUserController(userUseCase UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u *UserController) GetUsers(w http.ResponseWriter, r *http.Request, log logger.Logger) {
	log.Info("In GET Users Controller")
	users, err := u.userUseCase.GetUser()
	if err != nil {
		log.Error("Error getting the user in use case", logger.Err(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res, err := request.MarshalResponse(users, nil)
	if err != nil {
		log.Error("Error Marshalling", logger.Err(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
