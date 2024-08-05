package models

type PostTag struct {
	ID     int `json:"id" gorm:"column:ID;primaryKey;autoIncrement"`
	PostID int `json:"postID" gorm:"column:postID;not null"`
	TagID  int `json:"tagID" gorm:"column:tagID;not null"`
}

func (PostTag) TableName() string {
	return "PostTags"
}
