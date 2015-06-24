package main

import (
	"strconv"
)

func CreateUser(name string) (result string, err error) {
	_, err = Db.Query("INSERT INTO users (name) VALUES ($1)", name)
	result = "Created user with name:" + name
	return
}

func UpdateUserName(id int, name string) (result string, err error) {
	_, err = Db.Query("UPDATE users SET name=$1 WHERE id=$2", name, id)
	result = "Updated the user with id " + strconv.Itoa(id)
	return
}

func SelectAllUsers() (result string, err error) {

	rows, err := Db.Query("SELECT * FROM users")
	for rows.Next() {
		var in int
		var st string
		rows.Scan(&in, &st)
		result += st + "\n"
	}
	return
}

func DeleteUserById(id int) (result string, err error) {
	_, err = Db.Query("DELETE FROM users WHERE id = $1", id)
	result = "Delete user with id:" + strconv.Itoa(id)
	return
}

func SearchUserWithName(name string) (result string, err error) {
	rows, err := Db.Query("SELECT name FROM users WHERE name LIKE $1 || '%'", name)
	for rows.Next() {
		var st string
		rows.Scan(&st)
		result += st + "\n"
	}
	return
}

func GetUserFromDB(id int) (result string, err error) {
	rows, err := Db.Query("SELECT name FROM users WHERE id = $1", id)
	rows.Next()
	rows.Scan(&result)
	return
}
