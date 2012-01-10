package sansmagic

import (
	"mustache"
	//"io/ioutil"
	"path/filepath"
	"os"
	"log"
)

// ------------------------------------------------------------------
type TemplateMgr struct {
	Templates map[string]*mustache.Template
}

func NewTemplateMgr(paths []string) *TemplateMgr {
	tm := new(TemplateMgr)
	tm.Templates = make(map[string]*mustache.Template)
	
	errChan := make(chan os.Error)
	for _, p := range paths {				
		filepath.Walk(p, tm, errChan)
	}

	return tm
}

func (tm TemplateMgr) Get(path string) *mustache.Template {
	temp, ok := tm.Templates[path]
	if !ok {
		log.Fatal("Template: ", path, " not found")
	}
	return temp
}

func (tm *TemplateMgr) VisitDir(path string, f *os.FileInfo) bool {
	// what is this supposed to do? read the go docs on this method.
	return true
}

func (tm *TemplateMgr) VisitFile(path string, f *os.FileInfo) {
	temp, err := mustache.ParseFile(path)
	DieIf(err)

	if prev, ok := tm.Templates[path]; ok {
		log.Print(f)
		log.Fatal("overloading templates: ", prev)
	}
	tm.Templates[path] = temp	
}
