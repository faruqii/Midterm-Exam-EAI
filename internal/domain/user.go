package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// BeforeCreate generates a UUID for primary key before creating a new user
func (user *User) BeforeCreate(_ *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
