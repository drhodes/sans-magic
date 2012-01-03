package sansmagic

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Debug(x interface{}) {
	fmt.Printf("%v\n", x)
}

func LogIf(err os.Error) {
	if err != nil {
		log.Print(err)
	}
}

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
