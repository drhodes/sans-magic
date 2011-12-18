package nomagic


type Visitor interface {
	VisitTable(TableI)
	VisitField(FieldI)
}
