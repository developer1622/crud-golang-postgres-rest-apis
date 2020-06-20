package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "Postgres"
	DB_NAME     = "todo_db"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	sqlStatement := `INSERT INTO task VALUES(155452,'dffdf44444d','as44taxie');`
	_, err = db.Exec(sqlStatement)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
