package models

type RoleMenu struct {
	Id     int  `gorm:"type:integer" json:"id"`
	RoleID int  `gorm:"type:integer;column:id_role" json:"id_role"`
	MenuID int  `gorm:"type:integer;column:id_menu" json:"id_menu"`
	Role   Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Menu   Menu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RoleMenu) TableName() string {
	return "daftar_role_menu"
}
