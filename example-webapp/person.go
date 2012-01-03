package main

import (
	. "sansmagic"
	"launchpad.net/gobson/bson"
	//"launchpad.net/mgo"
)


// A person model --------------------------------------------------
type Person struct {
	Name string
	Email string
	Age int32
}

func (self *Person) GetByName(name string) {
	c := session.DB("sans").C("person")
	err := c.Find(bson.M{"name": name}).One(self)
	DieIf(err)	
}

func (self Person) Insert () {
	c := session.DB("sans").C("person")
	DieIf(c.Insert(self))
}
