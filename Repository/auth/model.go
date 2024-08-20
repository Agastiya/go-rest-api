package auth

import "go-rest-api/Dto"

type AuthRepository struct{}

type LoginData struct {
	Dto.Users
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Active      bool   `json:"active"`
	AccessToken string `json:"accessToken"`
}
