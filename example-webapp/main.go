package main

import (
	. "sansmagic"
    "http"
	"launchpad.net/mgo"
)

var (
	session, err = mgo.Mongo("localhost")
)

func main() {
	r := NewRouter([]Routable{
		new(Homepage),
		new(Homepage),
	})

	person := Person{ "Gopher", "gopher@example.com", 33 }
	person.Insert()

	DieIf(http.ListenAndServe(":8080", r.Mux()))
	// goto http://localhost:8080/Homepage/Gopher
}
