package main

import (
	"barcelonaZoo/api"
	"barcelonaZoo/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	// DBインスタンスの作成
	dbInstance := db.CreateSQLInstance()
	defer dbInstance.Close()

	r := api.GetRouter()
	r.Run()
}
