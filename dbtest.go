package main

import (
	//"bufio"
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"strconv"
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
	credentials, ferr := ioutil.ReadFile("config/dbconfig.config")
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

func CreateUser(name string) string {
	db.Query("INSERT INTO users (name) VALUES ($1)", name)
	return "Created user with name:" + name
}

func SelectAllUsers() string {

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	result := ""
	for rows.Next() {
		var in int
		var st string
		rows.Scan(&in, &st)
		result += st + "\n"
	}
	return result
}

func DeleteUserById(id int) string {
	db.Query("DELETE FROM users WHERE id = $1", id)
	return "Delete user with id:" + strconv.Itoa(id)
}

func SearchUserWithName(name string) string {
	rows, err := db.Query("SELECT name FROM users WHERE name LIKE $1 || '%'", name)
	if err != nil {
		log.Fatal(err)
	}
	result := ""
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return result
}

func GetUserFromDB(id int) string {
	rows, err := db.Query("SELECT name FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	var username string
	rows.Scan(&username)
	return username
}

func GivenGroupIdFindUsers(id int) string {

	query := "SELECT u.name FROM groups AS g INNER JOIN user_group_relations AS ugr ON g.id = ugr.group_id AND  g.id = $1 INNER JOIN users AS u ON u.id = ugr.user_id"

	rows, err := db.Query(query, id)

	if err != nil {
		log.Fatal(err)
	}
	result := ""
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return result
}

func GivenIdFindGroups(id int) string {

	query := "SELECT g.name FROM users AS u INNER JOIN user_group_relations AS ugr ON u.id = ugr.user_id AND u.id = $1 INNER JOIN groups AS g ON g.id = ugr.group_id"

	rows, err := db.Query(query, id)

	if err != nil {
		log.Fatal(err)
	}
	result := ""
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return result
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
