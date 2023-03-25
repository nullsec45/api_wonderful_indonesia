package models

type Kota struct {
	Id         int      `gorm:"type:integer;primaryKey" json:"id"`
	NamaKota   string   `gorm:"size:150" json:"nama_kota"`
	ProvinsiId int      `gorm:"type:integer;column:id_provinsi"  json:"id_provinsi" validate:"required"`
	Provinsi   Provinsi `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Kota) TableName() string {
	return "daftar_kota"
}
