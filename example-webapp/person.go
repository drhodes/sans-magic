package main

import . "sansmagic"

var Person = NewTable (
	"Person", // table name.
	
	Fields {
		"Name": String8{},
		"Email": String8{},
		"Age": Int8{},
		"Icon": LongBlob{},
	},
	
	Tables{},
	)







