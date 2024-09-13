package database

import "github.com/jmoiron/sqlx"

var DB *sqlx.DB

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func ConnectDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if the connection is established
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
