package main

import (
	//"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"strings"
)

var (
	db *sql.DB
)

//init - Init metod run before main.
func init() {
	/*
	*credentials.txt must have the following structure to work:
	*host_ip
	*port
	*user
	*password
	*dbname
	*sslmode*/
	credentials, ferr := ioutil.ReadFile("credentials.txt")
	check(ferr)
	config := strings.Split(string(credentials), "\n")

	dbConfig := "host=" + config[0] +
		" port=" + config[1] +
		" user=" + config[2] +
		" password=" + config[3] +
		" dbname=" + config[4] +
		" sslmode=" + config[5]

	var err error
	db, err = sql.Open("postgres", dbConfig)
	check(err)
}

//main - when go run dbtest.go is invoked this method is run.
func main() {
	selectAllUsers(db)
}

func selectAllUsers(db *sql.DB) {

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

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
