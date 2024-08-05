package models

type Category struct {
	CategoryID   int    `json:"categoryID" gorm:"column:categoryID;primaryKey;autoIncrement"`
	CategoryName string `json:"categoryName" gorm:"column:categoryName;type:varchar(45);unique"`
}

func (Category) TableName() string {
	return "Category"
}
