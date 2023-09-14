package domain

import (
	"github.com/eggysetiawan/gorahm-starterkit/dto"
	"github.com/eggysetiawan/gorahm-starterkit/errs"
)

type User struct{
	Name string
	Email string 
}

type UserRepository interface {
	Get()([]User,*errs.Exception)
	ById(string) (*User, *errs.Exception)
}


func (u User) ToDto() dto.UserResponse {
	return dto.UserResponse{
		Name:        u.Name,
		Email:       u.Email,
	}
}