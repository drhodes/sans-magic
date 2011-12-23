package main

import . "sansmagic"

var Person = NewTable (
	"Person",
	
	Fields {
		"Name": String8{},
		"Email": String8{},
		"Age": Int8{},
		"Icon": LongBlob{},
	},
	
	Tables{},
	)


