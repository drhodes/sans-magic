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

	err := http.ListenAndServe(":8080", r.Mux()) 
	if err != nil {
		log.Print(err)
	}
}
