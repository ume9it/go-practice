package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // ブランクインポート
)

// ConnectDB demonstrates connecting to the database.
func main() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if nil != err {
		fmt.Println("Error connecting to the database:", err)
		return
	} else {
		fmt.Println("Connected to the database successfully!")
	}

	err = db.Ping()
	if nil != err {
		fmt.Println("Error pinging the database:", err)
		return
	} else {
		fmt.Println("Ping to the database successful!")
	}
	defer db.Close()
}
