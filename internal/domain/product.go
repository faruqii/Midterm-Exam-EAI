package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Categories struct {
	ID   string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name string `json:"name"`
}

func (c *Categories) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return
}

type Product struct {
	ID          string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  string  `json:"category_id"`
	Category    Categories `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}
