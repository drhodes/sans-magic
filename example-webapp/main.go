package main

import . "nomagic"


func show(fld FieldI) {
	Debug(fld)
}


func main() {
	post :=	Post {
		Person { 
			String8{"Gertrude"},
			String8{"gertrude@gmail.com"},
			Int8{45},
			LongBlob{[]byte{1,7,65,4,3,4,6,7,8,89,67,56,54,4,4,65,6,4,5}},
		},
		String32{"I just don't like it, sir"},
	}
	
	//WalkTable(post, show)
	InsertTable(post)
}
