package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey;column:USER_ROW_ID"`
	UserName    string    `gorm:"column:USERNAME;not null"`
	Email       string    `gorm:"column:EMAIL;not null"`
	CreatedDate time.Time `gorm:"column:CREATED_DATE"`
	UpdatedDate time.Time `gorm:"column:UPDATED_DATE"`
}
