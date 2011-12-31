package main

import . "sansmagic"

var Post = NewTable (
	"Post",

	Fields { "Text": String32{} },		
	Tables { "Author": Person },
	)
