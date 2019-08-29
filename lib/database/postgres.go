package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Client *sql.DB
}

func OpenPostgres(databaseConnection string) (*PostgresDB, error) {
	DB, err := sql.Open("postgres", databaseConnection)
	if err != nil {
		return nil, err
	}
	if err = DB.Ping(); err != nil {
		DB.Close()
		return nil, err
	}
	pDB := PostgresDB{
		Client: DB,
	}
	return &pDB, nil
}
