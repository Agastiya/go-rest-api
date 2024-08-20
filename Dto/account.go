package Dto

import "time"

type CreateAccount struct {
	Users
	Name        string    `json:"name" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Phone       string    `json:"phone" validate:"required"`
	CreatedBy   string    `json:"createdBy" validate:"required"`
	CreatedTime time.Time `json:"createdTime" validate:"required"`
}

type UpdateAccountPassword struct {
	Users
	Id       int64  `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateAccountStatus struct {
	Users
	Id     int64 `json:"id" validate:"required"`
	Active *bool `json:"active" validate:"required"`
}
