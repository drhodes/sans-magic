package sansmagic

//import "http"

type Arg struct {
	Fields []string
}

func NewArg(... interface{}) Arg {
	return Arg{}
}


type Url struct {
	itemds []interface{}
}

func NewUrl(... interface{}) Url {
	return Url{}
}