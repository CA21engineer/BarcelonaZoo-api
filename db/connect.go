package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	DB *sql.DB
)

// CreateSQLInstance localMySQLへのコネクション作成
func CreateSQLInstance() *sql.DB {
	var err error

	log.Println("connectDB: mysql")
	dbuser := os.Getenv("MYSQL_USER")
	if dbuser == "" {
		dbuser = "root"
	}
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	if dbpassword == "" {
		dbpassword = "password"
	}
	dbhost := os.Getenv("MYSQL_HOST")
	if dbhost == "" {
		dbhost = "localhost"
	}
	dbname := "barcelona_zoo"

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbuser, dbpassword, dbhost, dbname)

	DB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	return DB
}
