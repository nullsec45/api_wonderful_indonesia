package models

type Provinsi struct {
	Id           int    `gorm:"type:integer;primaryKey"`
	NamaProvinsi string `gorm:"size:150"`
	IdRole       int    `gorm:"size:1;type:integer;default:5"`
}
