package domain

type Token struct {
	ID     string `json:"id" gorm:"primaryKey, type:uuid"`
	UserID string `json:"user_id" gorm:"type:uuid"`
	Token  string `json:"token"`
	Type   string `json:"type"`
	User   User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
