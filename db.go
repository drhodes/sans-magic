package nomagic

import (
	"fmt"
	"reflect"
	"os"
	"strings"
)

func Debug(x interface{}){
	fmt.Printf("%v\n", x)
}

func WalkTable(tbl TableI, vtor Visitor) {
	Debug("Inserting: " + tbl.TableName())
	for _, fld := range tbl.GetFields() {
		vtor.VisitField(fld)
	}
	for _, t := range tbl.GetForeignFields() {
		vtor.VisitTable(t)
		WalkTable(t, vtor)
	}		
}

type KeyId int64

func MagicTable(tbl TableI) map[string]string {
	reverseLookup := make(map[string]string)

	valOfTbl := reflect.ValueOf(tbl)	
	typeOfTbl := reflect.TypeOf(tbl)

	for i:=0; i < valOfTbl.NumField(); i++ {
		fldT := typeOfTbl.Field(i)
		field_name := fldT.Name
		field_tag := fldT.Tag
		
		fldV := valOfTbl.Field(i)

		if field_tag == "foreignkey" {
			reverseLookup[field_name] = "foreignkey"
		} else {
			reverseLookup[field_name] = fldV.Interface().(FieldI).SqlVal()
		}
	}

	return reverseLookup
}

func MagicLookup(m map[string]string, s string) (string, os.Error) {
	for k, v := range m {
		if v == s {
			return k, nil
		}	
	}
	return "not found", os.NewError("Could not lookup the field value for: " + s)
}

func InsertTable(tbl TableI) (KeyId, string) {
	Debug("Inserting: " + tbl.TableName())

	// We need to grab the name of the field from the table.  Not too much magic.
	revLookup := MagicTable(tbl)
	fld_pairs := make(map[string]string)

	for _, fld := range tbl.GetFields() {
		v := reflect.ValueOf(fld).Interface().(FieldI).SqlVal()
		field_name, err := MagicLookup(revLookup, v)
		if err == nil {
			revLookup[field_name] = "", false;
		} else {
			Debug(err)
			os.Exit(1)
		}
		fld_pairs[field_name] = v
	}

	foreign_pairs := make(map[string]string)
	
	recursive_qstring := ""
	fkeyID := -1
	
	for _, fld := range tbl.GetForeignFields() {
		t := reflect.ValueOf(fld).Interface().(TableI)
		fkeyID, qsr := InsertTable(fld)
		recursive_qstring += qsr 
		foreign_pairs[t.TableName()] = fmt.Sprintf("%d", fkeyID)
	}		

	QueryString := `INSERT INTO ` + tbl.TableName() + ` (%s) VALUES (%s);`

	comma_fields := []string{}
	comma_values := []string{}

	// fields (non foreign keys)
	for key, val := range fld_pairs {
		comma_fields = append(comma_fields, key)
		comma_values = append(comma_values, val) 
	}
	// foreign keys
	for key, val := range foreign_pairs {
		comma_fields = append(comma_fields, key)
		comma_values = append(comma_values, val) 
	}

	qs := fmt.Sprintf(QueryString, 
		strings.Join(comma_fields, ", "),
		strings.Join(comma_values, ", "))

	return KeyId(fkeyID), recursive_qstring + qs
}

