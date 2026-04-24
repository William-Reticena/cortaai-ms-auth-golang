package model

import "time"

type User struct {
	ID           int       `json:"id" gorm:"primaryKey;column:id"`
	Email        string    `json:"dsEmail" gorm:"uniqueIndex;column:ds_email"`
	PasswordHash string    `json:"-" gorm:"column:ds_password_hash"`
	Username     string    `json:"nmUser" gorm:"column:nm_user"`
	CreatedAt    time.Time `json:"dtCreated" gorm:"column:dt_created"`
	UpdatedAt    time.Time `json:"dtUpdated" gorm:"column:dt_updated"`
}

func (User) TableName() string {
	return "user"
}
