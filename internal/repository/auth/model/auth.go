package coremodel

import (
	"database/sql"
	"time"
)

type User struct {
	ID              int64        `db:"id"`
	Name            string       `db:"name"`
	Email           string       `db:"email"`
	Password        string       `db:"password"`
	PasswordConfirm string       `db:"password_confirm"`
	Role            string       `db:"role"`
	CreatedAt       time.Time    `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}
