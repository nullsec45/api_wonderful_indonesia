package models

type User struct {
	Id       int    `gorm:"type:integer;primaryKey"`
	Username string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	IdRole   int    `gorm:"size:2"`
}
