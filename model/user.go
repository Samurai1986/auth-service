package model

import (
	"github.com/google/uuid"
)

//user database model
type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MiddleName string `json:"middle_name"` 
}

//user json model
type UserDTO struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MiddleName string `json:"middle_name"` 
}

type LoginDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}
