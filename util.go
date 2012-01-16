package sansmagic

import (
	"fmt"
	"log"
	"os"
	"strings"
	"http"
)

func Debug(x interface{}) {
	fmt.Printf("%v\n", x)
}

func LogIf(err os.Error) os.Error {
	if err != nil {
		log.Print(err)
	} 
	return err
}

func DieIf(err os.Error, args ...interface{}) {
	if err != nil {
		log.Panic(err, args)
	}
}

func GetFormMap(req *http.Request) map[string]string {
	vals := req.FormValue("vals")
	fields := map[string]string{}

	if vals == "" {
		log.Print("Form values empty")
		return fields
	}

	for _, pair := range strings.Split(vals, "&") {
		ps := strings.Split(pair, "=")		
		fields[ps[0]] = ps[1]
	}
	return fields
}




/*
type MyString struct {
	S string
}

func (self *MyString) Write(p []uint8) (int, os.Error) {
	for _, b := range p {
		self.S += string(b)		
	}
	return 0, nil
}

func (self *MyString) Contains(s string) bool {
	return strings.Contains(self.S, s)
}
*/