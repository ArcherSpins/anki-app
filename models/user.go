package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type DBUser struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"unique"`
	Username     string    `json:"username" gorm:"unique"`
	PasswordHash string    `json:"-" gorm:"column:password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

type EditUser struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        *string   `json:"email"`
	Username     *string   `json:"username"`
	PasswordHash *string   `json:"-" gorm:"column:password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

func (User) TableName() string {
	return "accounts"
}

func (DBUser) TableName() string {
	return "accounts"
}

type UserResponse struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginData struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
