package app

import (
	context "context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

func CreateTable(pool *pgxpool.Pool) {
	sql := `CREATE TABLE IF NOT EXISTS params (
			article_id bigserial primary key,
			ema_min float8,
			ema_mid float8,
			ema_max float8
			)`
	_, err := pool.Exec(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
	}
}

func InsertData(pool *pgxpool.Pool, wg *sync.WaitGroup) {
	defer wg.Done()
	queryLst := make([][]any, 1001, 1001)
	for i := 1000; i <= 2000; i++ {
		mid := i + 1
		min := i + 10
		max := i + 100
		queryLst[i-1000] = []any{min, mid, max}
	}
	_, err := pool.CopyFrom(context.Background(), pgx.Identifier{"params"}, []string{"ema_min", "ema_mid", "ema_max"}, pgx.CopyFromRows(queryLst))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")
}
