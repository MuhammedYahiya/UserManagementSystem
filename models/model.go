package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
}

type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
	UserID   uint   `json:"-" gorm:"foreignKey:UserID"`
}
