package domain

type Role struct {
	ID   string `json:"id" gorm:"primaryKey, type:uuid"`
	Name string `json:"name"`
}
