package main

import . "nomagic"

type Person struct {
	Name String8
	Email String8
	Age Int8
	Icon LongBlob
}

func (p Person) GetFields() []FieldI {
	return []FieldI{
		p.Name,
		p.Email,
		p.Age,
		p.Icon,
	}
}

func (p Person) GetForeignFields() []TableI {
	return []TableI{}
}

func (p Person) TableName() string {
	return "Person"
}
