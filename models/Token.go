package models

import "time"

type Token struct {
	Id        int        `gorm:"type:integer;primaryKey"`
	UserId    int        `gorm:"type:integer;column:id_user" `
	MenuId    int        `gorm:"type:integer;column:id_menu"`
	token     string     `gorm:"size:255"`
	ExpiredAt *time.Time `gorm:"column:expired_at"`
	User      User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Menu      Menu       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Token) TableName() string {
	return "token_api"
}
