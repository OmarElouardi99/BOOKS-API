package data

import "time"

type user struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     token     `json:"token"`
}

type token struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"_"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
