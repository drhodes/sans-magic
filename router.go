package sansmagic

import (
	mux "gorilla.googlecode.com/hg/gorilla/mux"
	"http"
)

type Router struct {	
	Routables []Routable
}

type Routable interface {
	Route() string
	Handler() func (http.ResponseWriter, *http.Request) 
}

func NewRouter(rs []Routable) Router {
	return Router{ rs }
}

func (r Router) Server() *mux.Router {
	hdlr := new(mux.Router)
	for _, rs := range r.Routables {
		hdlr.HandleFunc(rs.Route(), rs.Handler())
	}
	return hdlr
}



