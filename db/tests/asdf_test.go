package sansmagic

import (
	"fmt"
	"launchpad.net/gobson/bson"
	"launchpad.net/mgo"
	"testing"
)

type Customer struct {
	Name  string
	Phone string
}

func TestMongo(t *testing.T) {
	t.Log("Looking for mongo daemon")
	session, err := mgo.Mongo("localhost")
	if err != nil {		
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Customer{"Ale", "+55 53 8116 9639"},
		&Customer{"Cla", "+55 53 8402 8510"})
	if err != nil {
		panic(err)
	}

	result := Customer{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)
}
