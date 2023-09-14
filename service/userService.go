package service

import (
	"github.com/eggysetiawan/gorahm-starterkit/domain"
	"github.com/eggysetiawan/gorahm-starterkit/dto"
	"github.com/eggysetiawan/gorahm-starterkit/errs"
)


type UserService interface {
	GetAllUser() ([]dto.UserResponse, *errs.Exception)
	GetUser(id string) (*dto.UserResponse, *errs.Exception)
}

type DefaultUserService struct{
	repo domain.UserRepository
}

func (s DefaultUserService) GetAllUser() ([]dto.UserResponse, *errs.Exception) {
// request phase
	// if status == "inactive" {
	// 	status = "0"
	// } else if status == "active" {
	// 	status = "1"
	// } else {
	// 	status = ""
	// }


	users, err := s.repo.Get()


// response phase
	if err != nil {
		return nil, err
	}
	response := make([]dto.UserResponse, 0)

	for _, u := range users {
		response = append(response, u.ToDto())
	}

	return response, nil
}

func (s DefaultUserService) GetUser(id string) (*dto.UserResponse, *errs.Exception) {
	u, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := u.ToDto()

	return &response, nil
}


func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}