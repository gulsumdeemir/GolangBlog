package models

import (
	"time"
)

type Tag struct {
	TagID   int       `json:"tagID" gorm:"column:tagID;primaryKey;autoIncrement"`
	TagName string    `json:"tagName" gorm:"column:tagName;type:varchar(45);not null"`
	Datee   time.Time `json:"datee" gorm:"column:datee;type:timestamp;default:CURRENT_TIMESTAMP"`
}


func (Tag) TableName() string {
	return "Tags"
}
