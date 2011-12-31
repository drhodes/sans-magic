package sansmagic

import (
	"os"
	"fmt"
	"strings"
	"log"
)

type KeyId int
type Tables map[string]Table

type Table struct {
	Name string
	Fields Fields // map[string]FieldI
	Tables Tables
	Id uint64
}

func (tbls Tables) Copy() Tables {
	m := make(map[string]Table)
	for key, val := range tbls {
		m[key] = val.Copy()
	}
	return m
}

func (tbl Table) Copy() Table {
	return Table {
		string(tbl.Name),
		tbl.Fields.Copy(),
		tbl.Tables.Copy(),
		uint64(tbl.Id),
	}
}

func (tbl Table) GetField(s string) FieldI {
	return tbl.Fields[s]
}

func (tbl Table) GetTable(s string) Table {
	return tbl.Tables[s]
}

func (tbl Table) UpdateField(s string, fld FieldI) Table {	
	t := tbl.Copy()
	t.Fields[s] = fld.Copy()
	return t
}

func (tbl Table) UpdateTable(s string, table Table) Table {	
	t := tbl.Copy()
	t.Tables[s] = table.Copy()
	return t
}

func NewTable(name string, flds Fields, tbls Tables) Table {
	return Table{name, flds, tbls, 0}
}

func WalkTable(tbl Table, vtor Visitor) {
	Debug("Inserting: " + tbl.Name)
	for _, fld := range tbl.Fields {
		vtor.VisitField(fld)
	}
	for _, t := range tbl.Tables {
		vtor.VisitTable(t)
		WalkTable(t, vtor)
	}		
}

func (tbl Table) Insert() (KeyId, string, os.Error) {
	Debug("Inserting: " + tbl.Name)
	QueryString := `INSERT INTO ` + tbl.Name + ` (%s) VALUES (%s);`

	recursive_qstring := ""
	fkeyID := -1

	comma_fields := []string{}
	comma_values := []string{}

	for fname, fld := range tbl.Fields {
		comma_fields = append(comma_fields, fname)
		comma_values = append(comma_values, fld.SqlVal()) 
	}
	
	for ftbl_name, ftbl := range tbl.Tables {
		fkeyID, qsr, _ := ftbl.Insert()
		recursive_qstring += qsr 
		keyIdTxt := fmt.Sprintf("%d", fkeyID)
		
		comma_fields = append(comma_fields, ftbl_name)
		comma_values = append(comma_values, keyIdTxt) 
	}
	
	qs := fmt.Sprintf(QueryString, 
		strings.Join(comma_fields, ", "),
		strings.Join(comma_values, ", "))
	
	return KeyId(fkeyID), recursive_qstring + qs, nil
}

func (tbl Table) Save() {
	_, instr, err := tbl.Insert()
	if err != nil {
		log.Print(instr)
	}	
}



/*
func Select(q Query) Table {	
}
*/