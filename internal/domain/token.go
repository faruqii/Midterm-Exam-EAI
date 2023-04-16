package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID     string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID string `json:"user_id"`
	Token  string `json:"token"`
	Type   string `json:"type"`
	User   User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return
}
