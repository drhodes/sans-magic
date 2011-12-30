package sansmagic

type Visitor interface {
	VisitTable(Table)
	VisitField(FieldI)
}
