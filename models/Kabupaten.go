package models

type Kabupaten struct {
	Id            int      `gorm:"type:integer;primaryKey" json:"id"`
	NamaKabupaten string   `gorm:"size:150;column:nama_kabupaten" json:"nama_kabupaten" validate:"required"`
	ProvinsiId    int      `gorm:"type:integer;column:id_provinsi"  json:"id_provinsi" validate:"required"`
	Provinsi      Provinsi `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Kabupaten) TableName() string {
	return "daftar_kabupaten"
}
