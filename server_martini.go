package main

import (
	"github.com/go-martini/martini"
	"strconv"
)

type Payload struct {
	data Data
}

type Data struct {
}

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World!"
	})
	m.Get("/api/1", apiV1)

	m.Group("/api/1/user", func(r martini.Router) {
		r.Get("/:id", GetUser)
		r.Get("/list/all", ListUsers)
		r.Post("/new/:name", NewUser)
		r.Put("/update/:id", UpdateUser)
		r.Delete("/delete/:id", DeleteUser)
		r.Get("/find/:name", FindUser)
	})

	m.Group("/api/1/group", func(r martini.Router) {
		r.Get("/user/list/all/:id", ShowGroups)
		r.Get("/list/users/all/:id", ShowUsersInGroup)
	})

	m.Run()
}

func apiV1() string {
	return "Hello API!"
}

func NewUser(params martini.Params) (int, string) {
	dbQuery := CreateUser(params["name"])
	return 200, dbQuery
}

func ListUsers() (int, string) {

	dbQuery := SelectAllUsers()
	return 200, dbQuery
}

func GetUser(params martini.Params) (int, string) {
	//auth
	//db uppslag baserat på id
	//returnera rimlig sak
	idParam := params["id"]
	intParam, _ := strconv.ParseInt(idParam, 10, 10)
	dbQuery := GetUserFromDB(int(intParam))
	return 200, dbQuery

}

func FindUser(params martini.Params) (int, string) {
	name := params["name"]
	dbQuery := SearchUserWithName(name)
	return 200, dbQuery
}

func UpdateUser(params martini.Params) (int, string) {
	return 501, "501 - Not Implemented"
}

func DeleteUser(params martini.Params) (int, string) {
	id, _ := strconv.Atoi(params["id"])
	id = int(id)
	dbQuery := DeleteUserById(id)
	return 200, dbQuery
}

func ShowGroups(params martini.Params) (int, string) {
	id, _ := strconv.Atoi(params["id"])
	id = int(id)
	dbQuery := GivenIdFindGroups(id)
	return 200, dbQuery
}

func ShowUsersInGroup(params martini.Params) (int, string) {
	id, _ := strconv.Atoi(params["id"])
	id = int(id)
	dbQuery := GivenGroupIdFindUsers(id)
	return 200, dbQuery
}
