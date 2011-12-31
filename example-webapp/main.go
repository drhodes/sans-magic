package main

import (
	. "sansmagic"
    "http"
    "log"
	//"io"
)
	
func main() {
	/*
	p := Person.UpdateField("Name", String8{"Derek"})
	log.Print(Person)
	log.Print(p)

	post := Post.UpdateTable("Author", p)
	a,b,c := post.Insert()
	log.Print(a, b, c)
	 */

	// Views.
	r := NewRouter([]Routable{
		NewHomepage(), 
		NewHomepage(),		
	})

	err := http.ListenAndServe(":8080", r.Mux()) 
	if err != nil {
		log.Print(err)
	}
}
