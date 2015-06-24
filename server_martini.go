package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

type User map[string]string

type Payload struct {
	Data Data
}

type Data struct {
	User User
}

func main() {
	m := martini.Classic()

	// map json encoder
	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Get("/", func() string {
		return "Hello World!"
	})
	m.Get("/api/1", apiV1)

	m.Group("/api/1/user", func(r martini.Router) {
		r.Get("/:id", GetUser)
		r.Get("/list/all", ListUsers)
		r.Post("/new/:name", NewUser)
		r.Put("/update/:id/:name", UpdateUser)
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
	dbQuery, _ := CreateUser(params["name"])
	return http.StatusOK, dbQuery
}

func ListUsers() (int, string) {

	dbQuery, _ := SelectAllUsers()
	return http.StatusOK, dbQuery
}

func GetUser(params martini.Params, w http.ResponseWriter) (int, string) {
	id, _ := strconv.Atoi(params["id"])
	dbQuery, _ := GetUserFromDB(int(id))

	response, err := buildUserResponse(int(id), dbQuery)
	Check(err)
	return http.StatusOK, string(response)
}

func FindUser(params martini.Params) (int, string) {
	name := params["name"]
	dbQuery, _ := SearchUserWithName(name)
	return http.StatusOK, dbQuery
}

func UpdateUser(params martini.Params) (int, string) {
	id, _ := strconv.Atoi(params["id"])
	id = int(id)
	name := params["name"]
	dbQuery, _ := UpdateUserName(id, name)
	return http.StatusOK, dbQuery
}

func DeleteUser(params martini.Params) (int, string) {
	id, _ := strconv.Atoi(params["id"])
	id = int(id)
	dbQuery, _ := DeleteUserById(id)
	return http.StatusOK, dbQuery
}

func ShowGroups(params martini.Params) (int, string) {
	id, _ := strconv.Atoi(params["id"])
	id = int(id)
	dbQuery, _ := GivenIdFindGroups(id)
	return http.StatusOK, dbQuery
}

func ShowUsersInGroup(params martini.Params) (int, string) {
	id, _ := strconv.Atoi(params["id"])
	id = int(id)
	dbQuery, _ := GivenGroupIdFindUsers(id)
	return http.StatusOK, dbQuery
}

func buildUserResponse(id int, name string) ([]byte, error) {
	user := make(map[string]string)
	user["id"] = strconv.Itoa(id)
	user["name"] = name

	d := Data{user}
	//p := Payload{d}
	return json.MarshalIndent(d, "", "  ")
}
