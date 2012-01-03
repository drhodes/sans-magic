package sansmagic

type Field struct {
	Name string     // column name in the schema
	Type SqlType    // column type in the schema
}

type Fields map[string]FieldI

func (self Fields) Copy() Fields {
	m := make(map[string]FieldI)
	for key, val := range self {
		m[key] = val.Copy()
	}
	return m
}

type FieldI interface {	
	MongoVal() string	
	Copy() FieldI
}


