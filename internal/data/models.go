package data

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const dbTimeOut = time.Second * 3

var db *sqlx.DB

func New(dbPool *sqlx.DB) Models {
	db = dbPool
	return Models{
		User:  User{},
		Token: Token{},
	}
}

type Models struct {
	User  User
	Token Token
}
type User struct {
	Id        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Password  []byte    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Token     Token     `json:"token"`
}

func (u *User) GetAll() ([]User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, email, first_name, last_name, password, created_at, updated_at from users ORDER BY last_name`
	users := []User{}
	err := db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, first_name, last_name, password, created_at, updated_at from users WHERE email = ?`
	user := User{}
	err := db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) GetById(id int) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, email, first_name, last_name, password, created_at, updated_at from users WHERE id = ?`
	user := User{}
	err := db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Update() error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `UPDATE users SET email = ?, first_name = ?, last_name = ?, updated_at=? WHERE id = ?`
	db.MustExecContext(ctx, query, u.Email, u.FirstName, u.LastName, time.Now(), u.Id)
	return nil
}

func (u *User) Delete(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `DELETE FROM users WHERE id = ?`
	db.MustExecContext(ctx, query, id)
	return nil
}

func (u *User) Add() error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	query := `INSERT INTO users (email, first_name, last_name, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	db.MustExecContext(ctx, query, u.Email, u.FirstName, u.LastName, u.Password, time.Now(), time.Now())

	return nil
}

func (u *User) ResetPassword(password string) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := `UPDATE users SET password = ?, updated_at = ? WHERE id = ?`
	db.MustExecContext(ctx, query, hashedPassword, time.Now(), u.Id)
	return nil
}

func (u *User) PasswordMatches(password string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

type Token struct {
	Id        int       `json:"id" db:"id"`
	UserId    int       `json:"user_id" db:"user_id"`
	Email     string    `json:"email" db:"email"`
	Token     string    `json:"token" db:"token"`
	TokenHash []byte    `json:"token_hash"`
	Expiry    time.Time `json:"expiry" db:"expiry"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (t *Token) GetByToken(plainText string) (*Token, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, use_id, email, token, token_hash, expiry, created_at, updated_at from tokens WHERE token = ?`
	token := Token{}
	err := db.GetContext(ctx, &token, query, plainText)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (t *Token) GetUserByToken(token Token) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	user := User{}
	query := `SELECT id, email, first_name, last_name, password, created_at, updated_at from users WHERE is = ?`
	err := db.GetContext(ctx, &user, query, token.UserId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
