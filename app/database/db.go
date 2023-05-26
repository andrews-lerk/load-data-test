package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func InitDB(url string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Print("Connection to database FAIL")
		log.Fatal(err)
	}
	log.Printf("Connection to database success: %s", url)
	return dbpool
}
