package sansmagic

import (
	"log"
	"http"
	"testing"
	"bytes"
	//"strings"
	"launchpad.net/gobson/bson"
	"launchpad.net/mgo"
	mux "gorilla.googlecode.com/hg/gorilla/mux"
)


var (
	session, err = mgo.Mongo("localhost")
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

// A post model ----------------------------------------------------
type Post struct {
	Text string
	Author Person
}

// A homepage view --------------------------------------------
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

// ------------------------------------------------------------------
func TestMain(t *testing.T) {
	DieIf(err)

	defer session.Close()

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

	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)

	person := Person{ "Derek", "asdf@asdf.com", 33 }
	person.Insert()
	person = Person{ "Fred", "asdf@a345sdf.com", 33 }
	person.Insert()

	person.GetByName("Fred")
	log.Print(person)
	person.GetByName("Derek")
	log.Print(person)

	buf := bytes.NewBuffer([]byte{})
	rsp.Write(buf)
	log.Print(buf.String())
}
