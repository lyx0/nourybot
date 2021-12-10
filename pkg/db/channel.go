package db

import (
	"database/sql"
)

type Channel struct {
	Id      int64
	Name    string
	Connect bool
}

type DB struct {
	db *sql.DB
}

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

// func (db *DB) Seed() error {
// 	data := Channel{
// 		Id:      31437432,
// 		Name:    "nouryxd",
// 		Connect: true,
// 	}
// }
