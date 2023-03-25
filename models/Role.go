package models

type Role struct {
	Id   int    `gorm:"type:integer;primaryKey" json:"id"`
	Role string `gorm:"size:150" json:"role"`
}

func (Role) TableName() string {
	return "daftar_role"
}
