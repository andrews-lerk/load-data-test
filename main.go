package main

import (
	"fmt"
	"go_example/app"
	"go_example/app/config"
	"go_example/app/database"
	"os"
	"sync"
	"time"
)

func main() {
	path := getPath()
	cfg := config.LoadConfig(path)
	pool := database.InitDB(cfg.DbUrl())
	defer pool.Close()
	app.CreateTable(pool)
	wg := sync.WaitGroup{}
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go app.InsertData(pool, &wg)
	}
	wg.Wait()
	fmt.Printf("Load finish in %v", time.Since(startTime))
}

func getPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}
