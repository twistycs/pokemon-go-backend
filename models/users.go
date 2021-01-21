package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey;column:USER_ROW_ID"  json:"-"`
	UserName    string    `gorm:"column:USERNAME;not null"`
	Password    string    `gorm:"column:PASSWORD;not null" json:"-"`
	Email       string    `gorm:"column:EMAIL;not null"`
	CreatedDate time.Time `gorm:"column:CREATED_DATE"`
	UpdatedDate time.Time `gorm:"column:UPDATED_DATE"`
}
