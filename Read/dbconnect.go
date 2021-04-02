package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/go-sql-driver/mysql.v1"
)

var db *sql.DB

// GetMongoDB function to return DB connection
func GetDBconn() *sql.DB {
	dbName := "userdb"
	fmt.Println("conn info:", dbName)
	db, err := sql.Open("mysql", "root:manish1234@tcp(localhost:3306)/userdb")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	return db
}
