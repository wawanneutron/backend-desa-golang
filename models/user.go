package models

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"PrimaryKey"`
	Name      string    `json:"name"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Email     string    `json:"Email" gorm:"unique;not null"`
	Password  string    `json:"-"` // tidak ditampilkan dalan response JSON
	Roles     []Role    `json:"roles" gorm:"many2many:user_roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
