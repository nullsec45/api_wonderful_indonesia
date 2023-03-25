package models

type User struct {
	Id       int    `gorm:"type:integer;primaryKey" json:"id"`
	Email    string `gorm:"size:255" json:"email"`
	Username string `gorm:"size:255" json:"username" validate:"required"`
	Password string `gorm:"size:255" json:"password" validate:"required"`
}

func (User) TableName() string {
	return "daftar_user"
}
