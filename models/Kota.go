package models

type Kota struct {
	Id         int    `gorm:"type:integer;primaryKey"`
	NamaKota   string `gorm:"size:150"`
	IdProvinsi int    `gorm:"type:integer;size:2"`
	IdRole     int    `gorm:"type:integer;size:1;default:5"`
}
