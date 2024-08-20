package auth

import (
	"go-rest-api/Config"
	"go-rest-api/Dto"
)

func (r AuthRepository) Login(params Dto.Login) (result LoginData, err error) {
	db := Config.DATABASE_MAIN.Get()
	resultData := db.Select("id, name, username, email, phone, active").Where("username = ? AND password = ? ", params.Username, params.Password).First(&result)
	err = resultData.Error
	return
}
