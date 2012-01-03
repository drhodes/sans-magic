package sansmagic

import (
	"http"
	"testing"
	mux "gorilla.googlecode.com/hg/gorilla/mux"
)

// A person model.
var Person = NewTable(
	"Person", // table name.

	Fields{
		"Name":  String{},
		"Email": String{},
		"Age":   Int32{},
		"Icon":  ByteArray{},
	},

	Tables{},
)

// A homepage view
type Homepage struct {
	Person Table
}

func NewHomepage() Homepage {
	return Homepage{Person}
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

	s := &String{}
	rsp.Write(s)

	if !s.Contains("Derek") {
		t.Fatal("ill formed response")
	}
}
