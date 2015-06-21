package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("postgres", "host=192.168.1.5 port=5432 user=postgres password=**** dbname=login_prototype sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var in int
		var st string
		rows.Scan(&in, &st)
		fmt.Printf("%s %d \n", st, in)
	}

}
