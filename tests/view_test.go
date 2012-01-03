package sansmagic

import (
	"log"
	"http"
	"testing"
	"bytes"
	"strings"
	"launchpad.net/gobson/bson"
	"launchpad.net/mgo"
	mux "gorilla.googlecode.com/hg/gorilla/mux"
)

// A person model --------------------------------------------------
type Person struct {
	Name string
	Email string
	Age int32
}

func (self Person) GetByName(session *mgo.Session, name string) Person {
	c := session.DB("test").C("people")
	result := new(Person)
	err := c.Find(bson.M{"name": "Derek"}).One(result)
	DieIf(err)	
	return *result
}

func (self Person) Insert (session *mgo.Session) {
	c := session.DB("test").C("people")
	DieIf(c.Insert(self))
}

// A post model ----------------------------------------------------
type Post struct {
	Text string
	Author Person
}

// A homepage view
type Homepage struct {
}

func NewHomepage() Homepage {
	return Homepage{}
}

func (hp Homepage) Route() string {
	return "/Homepage/{person:[A-Za-z0-9]+}"
}

func (hp Homepage) GetPerson(req *http.Request) []byte {
	p := mux.Vars(req)["person"]
	return []byte("{person:" + p + "}")
}

func (hp Homepage) Handler() func(rw http.ResponseWriter, request *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(hp.GetPerson(req))
	}
}

func TestMain(t *testing.T) {
	r := NewRouter([]Routable{
		NewHomepage(),
		NewHomepage(),
	})

	serve := func() {
		err := http.ListenAndServe(":8080", r.Mux())
		t.Fatal(err)
	}
	go serve()

	rsp, err := http.Get("http://localhost:8080/Homepage/Derek")
	LogIf(err)
	
	session, err := mgo.Mongo("localhost")
	if err != nil {		
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	person := Person{ "Derek", "asdf@asdf.com", 34 }
	person.Insert(session)

	p2 := Person{}
	log.Println(p2.GetByName(session, "Derek"))


	bs := []byte{}
	buf := bytes.NewBuffer(bs)
	rsp.Write(buf)

	if !strings.Contains(buf.String(), "Derek") {
		t.Fatal("ill formed response")
	}
}
