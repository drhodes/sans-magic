package sansmagic

import (
	//"fmt"
	//"launchpad.net/gobson/bson"
	//"launchpad.net/mgo"
	//"log"
	//"os"
	//"strings"
)

type KeyId int
type Tables map[string]Table

type Table struct {
	Name string
	Fields Fields // map[string]FieldI
	Tables Tables
	Id uint64
}







/*
type Customer struct {
	Name  string
	Phone string
}

func TestMongo45() {
	session, err := mgo.Mongo("localhost")
	if err != nil {		
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(
		&Customer{"Ale", "+55 53 8116 9639"},
		&Customer{"Cla", "+55 53 8402 8510"},
	)
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
*/

func (tbls Tables) Copy() Tables {
	m := make(map[string]Table)
	for key, val := range tbls {
		m[key] = val.Copy()
	}
	return m
}

func (tbl Table) Copy() Table {
	return Table {
		string(tbl.Name),
		tbl.Fields.Copy(),
		tbl.Tables.Copy(),
		uint64(tbl.Id),
	}
}

func (tbl Table) GetField(s string) FieldI {
	return tbl.Fields[s]
}

func (tbl Table) GetTable(s string) Table {
	return tbl.Tables[s]
}

func (tbl Table) UpdateField(s string, fld FieldI) Table {	
	t := tbl.Copy()
	t.Fields[s] = fld.Copy()
	return t
}

func (tbl Table) UpdateTable(s string, table Table) Table {	
	t := tbl.Copy()
	t.Tables[s] = table.Copy()
	return t
}

func NewTable(name string, flds Fields, tbls Tables) Table {
	return Table{name, flds, tbls, 0}
}
