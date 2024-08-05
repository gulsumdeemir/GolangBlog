package models

import (
	"time"
)

type Post struct {
	PostID          int       `json:"postID" gorm:"column:postID;primaryKey;autoIncrement"`
	PostTitle       string    `json:"postTitle" gorm:"column:postTitle;type:varchar(45);not null"`
	PostDescription string    `json:"postDescription" gorm:"column:postDescription;type:varchar(300);not null"`
	Datee           time.Time `json:"datee" gorm:"column:datee;type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID          int       `json:"userID" gorm:"column:userID;not null"`
	CategoryID      *int      `json:"categoryID" gorm:"column:categoryID"`

}


func (Post) TableName() string {
	return "Posts"
}
