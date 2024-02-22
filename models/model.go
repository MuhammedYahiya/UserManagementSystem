package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
}

type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `gorm:"not null"`
}
