package main

import (
	mux "gorilla.googlecode.com/hg/gorilla/mux"
	"http"
	"fmt"
)

// A homepage view --------------------------------------------
type Homepage struct {}

func (hp Homepage) Route() string {
	return "/Homepage/{person:[A-Za-z0-9]+}"
}

func (hp Homepage) GetPerson(req *http.Request) []byte {
	p := mux.Vars(req)["person"]
	
	person := Person{}
	person.GetByName(p)

	return []byte(fmt.Sprintf(`{person:"%s"}`, person.Email))
}

func (hp Homepage) Handler() func(rw http.ResponseWriter, request *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(hp.GetPerson(req))
	}
}
