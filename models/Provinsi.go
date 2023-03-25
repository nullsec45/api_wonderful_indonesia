package models

type Provinsi struct {
	Id           int    `gorm:"type:integer;primaryKey" json:"idc"`
	NamaProvinsi string `json:"nama_provinsi" gorm:"size:150;column:nama_provinsi" validate:"required"`
}

func (Provinsi) TableName() string {
	return "daftar_provinsi"
}
