package models

type PakaianAdat struct {
	Id            int      `gorm:"type:integer;primaryKey" json:"id"`
	NamaPakaian   string   `gorm:"size:100;column:nama_pakaian" json:"nama_pakaian" validate:"required""`
	GambarPakaian string   `gorm:"size:255;column:gambar_pakaian" json:"nama_gambar" validate:"required"`
	ProvinsiId    int      `gorm:"type:integer;column:id_provinsi"  json:"id_provinsi" validate:"required"`
	Provinsi      Provinsi `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (PakaianAdat) TableName() string {
	return "daftar_pakaian_adat"
}
