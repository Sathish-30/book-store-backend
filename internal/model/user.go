package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	EmailId  string `json:"email_id"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}
