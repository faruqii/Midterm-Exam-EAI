package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserBalance struct {
	ID      string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance" default:"0"`
	User    User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

func (t *UserBalance) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return
}
