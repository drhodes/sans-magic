include $(GOROOT)/src/Make.inc

TARG=sansmagic

GOFILES=\
	template.go\
	model.go\
	util.go\
	table.go\
	fields.go\
	view.go\
	visitor.go\
	router.go\
	db/*.go\

include $(GOROOT)/src/Make.pkg
