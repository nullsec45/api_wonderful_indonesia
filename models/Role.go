package models

type Role struct {
	Id   int    `gorm:"type:integer;primaryKey"`
	Role string `gorm:"size:150"`
	Kode string `gorm:"type:char;size:2"`
}
