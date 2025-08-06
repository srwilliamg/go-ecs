package controller

import (
	"net/http"
	customError "srwilliamg/app/v1/internal/application/custom-error"
	"srwilliamg/app/v1/internal/application/dto"
	"srwilliamg/app/v1/internal/application/request"
	"srwilliamg/app/v1/internal/interfaces/logger"
)

type UserUseCase interface {
	GetUser() ([]dto.User, error)
	CreateUser(dto.User) error
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
	response, err := u.userUseCase.GetUser()
	if err != nil {
		cerr := customError.NewCustomError("Error fetching users", nil)
		request.PrepareResponse(&w, cerr, http.StatusInternalServerError, log)
	}

	request.PrepareResponse(&w, response, http.StatusOK, log)
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request, log logger.Logger) {
	log.Info("In POST User Controller")
	var user dto.User
	if err := request.DecodeBody(r.Body, &user); err != nil {
		cerr := customError.NewCustomError("Error decoding JSON body", nil)
		request.PrepareResponse(&w, cerr, http.StatusBadRequest, log)
		return
	}

	err := u.userUseCase.CreateUser(user)
	if err != nil {
		cerr := customError.NewCustomError("Error creating user", []string{err.Error()})
		request.PrepareResponse(&w, cerr, http.StatusInternalServerError, log)
		return
	}

	request.PrepareResponse(&w, user, http.StatusCreated, log)
}
