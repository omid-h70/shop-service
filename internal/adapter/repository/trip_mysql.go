package repository

import (
	"database/sql"
)

type TripRepositoryMySqlDB struct {
	client *sql.DB
}
