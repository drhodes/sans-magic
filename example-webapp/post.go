package main

import . "nomagic"

type Post struct {
	Author Person 
	Text String32
}

func (p Post) GetFields() []FieldI {
	return []FieldI{p.Text}
}

func (p Post) GetRecFields() []TableI {
	return []TableI{p.Author}
}

func (p Post) TableName() string {
	return "Post"
}