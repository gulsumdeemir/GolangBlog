package models

import (
	"time"
)

type Save struct {
	SaveID     int       `json:"saveID" gorm:"column:saveID;primaryKey;autoIncrement"`
	UserID     int       `json:"userID" gorm:"column:userID;not null"`
	PostID     int       `json:"postID" gorm:"column:postID;not null"`
	CategoryID int       `json:"categoryID" gorm:"column:categoryID;not null"`
	Datee      time.Time `json:"datee" gorm:"column:datee;type:timestamp;default:CURRENT_TIMESTAMP"`

}


func (Save) TableName() string {
	return "Save"
}
