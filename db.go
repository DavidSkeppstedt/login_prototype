package main

import (
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
	var err error
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

	db, err = sql.Open("postgres", dbConfig)
	check(err)
}

func CreateUser(name string) (result string, err error) {
	_, err = db.Query("INSERT INTO users (name) VALUES ($1)", name)
	result = "Created user with name:" + name
	return
}

func UpdateUserName(id int, name string) (result string, err error) {
	_, err = db.Query("UPDATE users SET name=$1 WHERE id=$2", name, id)
	result = "Updated the user with id " + strconv.Itoa(id)
	return
}

func SelectAllUsers() (result string, err error) {

	rows, err := db.Query("SELECT * FROM users")
	for rows.Next() {
		var in int
		var st string
		rows.Scan(&in, &st)
		result += st + "\n"
	}
	return
}

func DeleteUserById(id int) (result string, err error) {
	_, err = db.Query("DELETE FROM users WHERE id = $1", id)
	result = "Delete user with id:" + strconv.Itoa(id)
	return
}

func SearchUserWithName(name string) (result string, err error) {
	rows, err := db.Query("SELECT name FROM users WHERE name LIKE $1 || '%'", name)
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return
}

func GetUserFromDB(id int) (result string, err error) {
	rows, err := db.Query("SELECT name FROM users WHERE id = $1", id)
	rows.Next()
	rows.Scan(&result)
	return
}

func GivenGroupIdFindUsers(id int) (result string, err error) {

	query := "SELECT u.name FROM groups AS g INNER JOIN user_group_relations AS ugr ON g.id = ugr.group_id AND  g.id = $1 INNER JOIN users AS u ON u.id = ugr.user_id"

	rows, err := db.Query(query, id)
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return
}

func GivenIdFindGroups(id int) (result string, err error) {

	query := "SELECT g.name FROM users AS u INNER JOIN user_group_relations AS ugr ON u.id = ugr.user_id AND u.id = $1 INNER JOIN groups AS g ON g.id = ugr.group_id"

	rows, err := db.Query(query, id)
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
