include $(GOROOT)/src/Make.inc

TARG=sansmagic

GOFILES=\
	model.go\
	util.go\
	table.go\
	fields.go\
	view.go\
	visitor.go\
	url.go\
	router.go\
	db/*.go\

include $(GOROOT)/src/Make.pkg
