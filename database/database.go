package database

import (
	"database/sql"
	"os"

	"github.com/chamick7/short-service/sqlc"
	_ "github.com/lib/pq"
)

func Init() *sqlc.Queries {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return sqlc.New(db)
}
