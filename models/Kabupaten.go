package models

type Kabupaten struct {
	Id            int    `gorm:"type:integer;primaryKey"`
	NamaKabupaten string `gorm:"size:150"`
	IdProvinsi    int    `gorm:"type:integer;size:2"`
	IdRole        int    `gorm:"size:1;type:integer;default:5"`
}
