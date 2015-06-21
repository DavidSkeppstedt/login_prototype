package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World!"
	})
	m.Get("/api/v1", apiV1)
	m.Run()
}

func apiV1() string {
	return "Hello API!"
}
