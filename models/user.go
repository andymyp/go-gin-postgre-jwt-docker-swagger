package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);" json:"name" validate:"required"`
	Email     string    `gorm:"type:varchar(50);unique;" json:"email" validate:"required,email"`
	Password  string    `gorm:"type:text" json:"password" validate:"required,min=6"`
	Token     string    `gorm:"type:text" json:"token"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return nil
}

type Claims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}

type LoginResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type UserRegister struct {
	Name     string `gorm:"type:varchar(100);" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(50);unique;" json:"email" validate:"required,email"`
	Password string `gorm:"type:text" json:"password" validate:"required,min=6"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
