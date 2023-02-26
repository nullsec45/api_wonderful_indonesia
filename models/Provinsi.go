package models

type Provinsi struct {
	Id           int    `gorm:"type:integer;primaryKey"`
	NamaProvinsi string `gorm:"size:150"`
	IdRole       int    `gorm:"type:integer;size:1;default:5"`
}
