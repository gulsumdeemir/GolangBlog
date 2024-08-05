package models

import (
	"time"
)

type Comment struct {
	CommentID int       `json:"commentID" gorm:"column:commentID;primaryKey;autoIncrement"`
	Commentt  string    `json:"commentt" gorm:"column:commentt;type:varchar(50);not null"`
	Datee     time.Time `json:"datee" gorm:"column:datee;type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID    int       `json:"userID" gorm:"column:userID;not null"`
	PostID    *int      `json:"postID" gorm:"column:postID"`

}


func (Comment) TableName() string {
	return "Comments"
}
