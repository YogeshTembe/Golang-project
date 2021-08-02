package model

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber int       `json:"phoneNumber"`
	IsActive    bool      `json:"isActive"`
}
