include $(GOROOT)/src/Make.inc

TARG=nomagic

GOFILES=\
	model.go\
	db.go\
	visitor.go\

include $(GOROOT)/src/Make.pkg
