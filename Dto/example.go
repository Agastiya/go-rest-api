package Dto

import "time"

func (Example) TableName() string {
	return "example"
}

type Example struct {
	Id          int64      `json:"id" gorm:"primaryKey"`
	Status      string     `json:"status"`
	Date        string     `json:"date"`
	Reason      string     `json:"reason"`
	CreatedBy   string     `json:"createdBy"`
	CreatedTime *time.Time `json:"createdTime"`
}
