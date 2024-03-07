package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int64
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

type UserCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
}

type UserUpdate struct {
	ID    int64
	Name  *string
	Email *string
	Role  *string
}
