package main

import (
	"http"
	. "sansmagic"
	"log"
)

// this is the shape of the json response.
// 

type Homepage struct {	
	Author Table
	Posts []Table
} 

/*
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
*/

type Foo struct {
}

func (f Foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print(w)
	log.Print(r)
}

func (hp Homepage) Route() Foo {
	return Foo{}
}


func (hp Homepage) Response(... string) string {
	return "asdf"
}

