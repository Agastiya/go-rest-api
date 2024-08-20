package auth

import (
	"go-rest-api/Config"
	"go-rest-api/Dto"
)

func (r AuthRepository) ValidateAccountData(field string, value string) (result bool, err error) {
	db := Config.DATABASE_MAIN.Get()
	if field == "email" {
		db.Raw("select Exists(select * from users Where email = ?)", value).Scan(&result)
	} else if field == "username" {
		db.Raw("select Exists(select * from users Where username = ?)", value).Scan(&result)
	} else if field == "phone" {
		db.Raw("select Exists(select * from users Where phone = ?)", value).Scan(&result)
	}
	return
}

func (r AuthRepository) CreateAccount(params Dto.CreateAccount) error {
	db := Config.DATABASE_MAIN.Get()
	result := db.Create(&params)
	return result.Error
}

func (r AuthRepository) UpdateAccountPassword(params Dto.UpdateAccountPassword) error {
	db := Config.DATABASE_MAIN.Get()
	result := db.Model(&Dto.UpdateAccountPassword{}).Where("id = ?", params.Id).Update("password", params.Password)
	return result.Error
}

func (r AuthRepository) UpdateAccountStatus(params Dto.UpdateAccountStatus) error {
	db := Config.DATABASE_MAIN.Get()
	result := db.Model(&Dto.UpdateAccountStatus{}).Where("id = ?", params.Id).Update("active", params.Active)
	return result.Error
}
