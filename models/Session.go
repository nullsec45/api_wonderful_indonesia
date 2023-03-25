package models

import "time"

type Session struct {
	Id        int        `gorm:"type:integer"`
	TokenId   int        `gorm:"type:integer;column:id_token"`
	ExpiredAt *time.Time `gorm:"column:expired_at"`
	Token     Token      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Session) TableName() string {
	return "session_token"
}
