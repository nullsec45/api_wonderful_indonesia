package models

type UserRole struct {
	Id     int  `gorm:"type:integer"`
	UserId int  `gorm:"type:integer;column:id_user"`
	RoleId int  `gorm:"type:integer;column:id_role"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role   Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (UserRole) TableName() string {
	return "daftar_user_role"
}
