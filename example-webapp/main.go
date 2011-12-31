package main

import (
	. "sansmagic"
    "http"
    "log"
)
	
func main() {
	r := NewRouter([]Routable{
		NewHomepage(), 
		NewHomepage(),		
	})
	
	log.Print(r)
	
	http.ListenAndServe(":8080", r.Mux()) 
}
