package db

import "database/sql"

const DSN = "postgres://username:password@localhost:5432/dbname?sslmode=disable"

func Connect() (*sql.DB, error) {
	return sql.Open("postgres", DSN)
}
