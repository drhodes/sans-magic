include $(GOROOT)/src/Make.inc

TARG=nomagic

GOFILES=\
	model.go\
	db.go\

include $(GOROOT)/src/Make.pkg
