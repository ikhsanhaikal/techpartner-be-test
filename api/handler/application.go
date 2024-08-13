package handler

import "database/sql"

type Application struct {
	DB        *sql.DB
	SecretKey []byte
}
