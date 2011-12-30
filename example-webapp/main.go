package main

import (
	. "sansmagic"
    "http"
    "log"
	"io"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
	
func main() {
	p := Person.UpdateField("Name", String8{"Derek"})
	log.Print(Person)
	log.Print(p)

	post := Post.UpdateTable("Author", p)
	a,b,c := post.Insert()
	log.Print(a, b, c)

	r := NewRouter([]Routable{
		NewHomepage(),
		NewHomepage(),
	})
	log.Print(r)
	
	http.ListenAndServe(":8080", r.Server()) 
}
