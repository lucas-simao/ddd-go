package repository

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New() Repository {
	dataSource, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Panic("Error to get DATABASE_URL")
	}

	newDb, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Panic(err)
		return &repository{}
	}

	return &repository{
		db: newDb,
	}
}
