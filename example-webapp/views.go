package main

import (
	. "sansmagic"
	"http"
	"log"
)

// A request comes down and a 


type Homepage struct {	
	Post Table	
} 

func NewHomepage() Homepage {
	return Homepage { Post }
}

func (hp Homepage) Route() string {
	return "/Homepage/{postid:[0-9]+}"
}


func (hp Homepage) Handler() func(rw http.ResponseWriter, request *http.Request) {
	return func(rw http.ResponseWriter, request *http.Request) {
		log.Print(rw)
		log.Print(request)
	}
}

