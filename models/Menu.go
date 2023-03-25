package models

type Menu struct {
	Id        int    `gorm:"type:integer;primaryKey" json:"id"`
	Menu      string `gorm:"type:string;column:nama_menu;size:255" json:"nama_menu"`
	kode_menu string `gorm:"size:45;column:kode_menu;" json:"id_role"`
}

func (Menu) TableName() string {
	return "daftar_menu"
}
