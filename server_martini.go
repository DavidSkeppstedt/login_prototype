package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World!"
	})
	m.Get("/api/v1", apiV1)

	m.Group("/api/v1/user", func(r martini.Router) {
		r.Get("/:id", GetUser)
		r.Get("/list/all", ListUsers)
		r.Post("/new/", NewUser)
		r.Put("/update/:id", UpdateUser)
		r.Delete("/delete/:id", DeleteUser)
	})

	m.Run()
}

func apiV1() string {
	return "Hello API!"
}

func NewUser(params martini.Params) (int, string) {
	return 200, "Added a new user!"
}

func ListUsers() (int, string) {
	return 501, "501 - Not Implemented"
}

func GetUser(params martini.Params) (int, string) {
	return 200, "{user{id:" + params["id"] + ",name:'David'}}"
}

func UpdateUser(params martini.Params) (int, string) {
	return 501, "501 - Not Implemented"
}

func DeleteUser(params martini.Params) (int, string) {
	return 501, "501 - Not Implemented"
}
