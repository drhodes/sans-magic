include $(GOROOT)/src/Make.inc

TARG=sansmagic

GOFILES=\
	model.go\
	util.go\
	db.go\
	table.go\
	fields.go\
	view.go\
	visitor.go\
	url.go\

include $(GOROOT)/src/Make.pkg
