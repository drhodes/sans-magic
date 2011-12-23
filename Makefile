include $(GOROOT)/src/Make.inc

TARG=sansmagic

GOFILES=\
	model.go\
	util.go\
	db.go\
	table.go\
	fields.go\
	routes.go\
	visitor.go\

include $(GOROOT)/src/Make.pkg
