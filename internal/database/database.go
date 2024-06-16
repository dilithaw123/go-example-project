package database

import "database/sql"

func New() *sql.DB {
	return &sql.DB{}
}
