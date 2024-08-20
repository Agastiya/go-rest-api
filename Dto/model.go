package Dto

type Users struct{}

func (Users) TableName() string {
	return "users"
}
