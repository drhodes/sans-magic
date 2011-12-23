package main

import "log"
import . "sansmagic"


func main() {
	p := Person.UpdateField("Name", String8{"Derek"})
	log.Print(Person)
	log.Print(p)

	post := Post.UpdateTable("Author", p)
	a,b,c := post.Insert()
	log.Print(a)
	log.Print(b)
	log.Print(c)


	


	//temp := A { { 4 } }
	//_, qs := InsertTable(post)
}




