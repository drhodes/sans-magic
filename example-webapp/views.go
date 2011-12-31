package main

import (
	. "sansmagic"
	mux "gorilla.googlecode.com/hg/gorilla/mux"
	"http"
	"log"
)
/*
 p := Person.UpdateField("Name", String8{"Derek"})
 log.Print(Person)
 log.Print(p)
 
 post := Post.UpdateTable("Author", p)
 a,b,c := post.Insert()
 log.Print(a, b, c)
 */

// Views.

type Homepage struct {	
	Person Table
} 

func NewHomepage() Homepage {
	return Homepage { Person }
}

func (hp Homepage) Route() string {
	return "/Homepage/{person:[A-Za-z0-9]+}"
}

func (hp Homepage) GetPerson(req *http.Request) []byte {
	p := mux.Vars(req)["person"]
	return []byte("{person: " + p + "}")
}

func (hp Homepage) Handler() func(rw http.ResponseWriter, request *http.Request) {	
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Print("Goodie")
		rw.Write(hp.GetPerson(req))
	}
}
