package sansmagic

import (
	"time"
	//"fmt"
)

func quote(s string) string { return `"` + s + `"` }

// ------------------------------------------------------------------
type String struct {
	Goval string
}


func (self String) MongoVal() string {
	return quote(string(self.Goval))
}

func (self String) Copy() FieldI {
 	return String{self.Goval}
}

// ------------------------------------------------------------------
type Int32 struct {	
	Goval int32
}


func (self Int32) Copy() FieldI {
 	return Int32{self.Goval}
}

func (self Int32) MongoVal() string {
	return string(self.Goval)
}

// ------------------------------------------------------------------
type Int64 struct {
	Goval int64
}
// ------------------------------------------------------------------
type Float64 struct {
	Goval float64
}
// ------------------------------------------------------------------
type Date struct {
	Goval time.Time
}
// ------------------------------------------------------------------
type ByteArray struct {
	Goval []FieldI
}
// ------------------------------------------------------------------
type Bool struct {
	Goval bool
}
// ------------------------------------------------------------------
type Null struct {
}

/*
type Object {
	Goval 
}
type Array {
	Goval
}
*/