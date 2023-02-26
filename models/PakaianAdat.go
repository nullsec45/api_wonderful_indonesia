package models

type PakaianAdat struct {
	Id            int    `gorm:"type:integer;primaryKey"`
	NamaPakaian   string `gorm:"size:150"`
	GambarPakaian string `gorm:"size:100"`
	IdDaerah      int    `gorm:"type:integer;size:3"`
	IdRole        int    `gorm:"type:integer;size:1;default:1"`
}
