package entity

import (
	"time"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Version   int        `json:"version"`
}

type RegisterPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
