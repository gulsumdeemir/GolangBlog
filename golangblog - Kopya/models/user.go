package models

import (
	"time"
)

type User struct {
    UserID       int         `json:"userID" gorm:"column:userID;primaryKey;autoIncrement"`
    UserName     string      `json:"userName" gorm:"column:userName;type:varchar(45);unique"`
    UserPassword string      `json:"userPassword" gorm:"column:userPassword;type:varchar(45)"`
    EMail        string      `json:"email" gorm:"column:eMail;type:varchar(45);unique"`
    Datee        *time.Time  `json:"date" gorm:"column:datee"`
}

