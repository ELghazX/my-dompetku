package user

import "database/sql"

type entity struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	Password  string       `db:"password"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
