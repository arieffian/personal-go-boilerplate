package models

type User struct {
	ID   string `json:"id" gorm:"primaryKey,column:id"`
	Name string `json:"name" goem:"column:name"`
}
