package main

import (
	. "sansmagic"
	mux "gorilla.googlecode.com/hg/gorilla/mux"
	"http"
	"log"
)


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
