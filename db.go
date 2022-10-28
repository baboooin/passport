package main

import (
	"database/sql"
_   "github.com/lib/pq"
)

func ConnectDatabase()(*sql.DB, error) {
	db, err := sql.Open("postgres",POSTGRES_URL)
	return db,err
}


func InitDatabase()(*sql.DB, error) {
	db, err := sql.Open("postgres",POSTGRES_URL)
	return db,err
}


