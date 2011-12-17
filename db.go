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

func WalkTable(tbl TableI, fn func(FieldI) ) {
	Debug("Inserting: " + tbl.TableName())
	for _, fld := range tbl.GetFields() {
		fn(fld)
	}
	for _, t := range tbl.GetRecFields() {
		WalkTable(t, fn)
	}		
}

type KeyId uint64

func isForeign(field_type reflect.Type) bool {
	Debug(field_type.Kind())
	return !strings.HasPrefix(field_type.PkgPath(), "nomagic")
}

func MagicTable(tbl TableI) map[string]string {
	reverseLookup := make(map[string]string)

	valOfTbl := reflect.ValueOf(tbl)	
	typeOfTbl := reflect.TypeOf(tbl)

	for i:=0; i < valOfTbl.NumField(); i++ {
		field_type := reflect.TypeOf(tbl).Field(i).Type;
		field_name := typeOfTbl.Field(i).Name

		if isForeign(field_type) {
			reverseLookup[field_name] = "foreignkey"
		} else {
			reverseLookup[field_name] = valOfTbl.Field(i).Interface().(FieldI).SqlVal()
		}
	}

	return reverseLookup
}

func MagicLookup(m map[string]string , s string) (string, os.Error) {
	for k, v := range m {
		if v == s {
			return k, nil
		}	
	}
	return "not found", os.NewError("Could not lookup the field value for: " + s)
}


func InsertTable(tbl TableI) KeyId {
	Debug("Inserting: " + tbl.TableName())

	Query := []string{}
	QueryString := `INSERT INTO ` + tbl.TableName() + ` (%s) VALUES (%s);`
	Debug(Query)

	// We need to grab the name of the field from the table.  Not too much magic.
	revLookup := MagicTable(tbl)
	Debug(revLookup)

	columns := []string{}
	values := []string{}	

	for _, fld := range tbl.GetFields() {
		v := reflect.ValueOf(fld).Interface().(FieldI).SqlVal()
		field_name, err := MagicLookup(revLookup, v)
		if err == nil {
			revLookup[field_name] = "", false;
		} else {
			Debug(err)
			os.Exit(1)
		}
		columns = append(columns, field_name)
		values = append(values, v)
		//Query = append(Query, fmt.Sprintf( QueryString, "asdf" , fld.SqlVal() )
	}

	Debug(columns)
	Debug(values)
	// need to strings.join(columns, ",")
	// need to strings.join(values, ",")
	// then get the keys from the InsertTable for any foreign keys

	for _, t := range tbl.GetRecFields() {
		InsertTable(t)
	}		

	Debug("QueryString: " + QueryString)
	return 666
}

