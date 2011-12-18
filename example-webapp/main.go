package main

import . "nomagic"

func show(fld FieldI) {
	Debug(fld)
}


func main() {
	post :=	Post {
		Person { 
			"Gary Gopher",
			"gary@gophertech.com",
			45,
			[]byte{1,7,65,4,3,4,6,7,8,89,67,56,54,4,4,65,6,4,5},
		},
		"No sir, I don't like it",
	}
	
	_, qs := InsertTable(post)
	Debug(qs)
}
