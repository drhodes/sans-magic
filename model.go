package nomagic

import (
	"time";
	"fmt";
)

func quote(s string) string { return `"` + s + `"` }

// Numeric Types ------------------------------------------------------------------
type Int8 struct {
	GoValue int8
}

func (self Int8) SqlType() string { 
	return "TINYINT" 
}

func (self Int8) SqlVal() string {
	return quote(fmt.Sprintf("%d", self.GoValue))
}

// ------------------------------------------------------------------
type Int16 struct {
	GoValue int16
}
type Int32 struct {
	GoValue int32
}
type Int64 struct {
	GoValue int64
}
type Uint8 struct {
	GoValue uint8
}
type Uint16 struct {
	GoValue uint16
}
type Uint32 struct {
	GoValue uint32
}
type Uint64 struct {
	GoValue uint64
}
type Float32 struct {
	GoValue float32
}
type Float64 struct {
	GoValue float64
}
type Bool struct {
	GoValue bool
}
type String struct {
	GoValue string 
}

// ------------------------------------------------------------------
type String8 struct {
	GoValue string 
}

func (self String8) SqlType() string { 
	return "TINYTEXT" 
}

func (self String8) SqlVal() string {
	return quote(self.GoValue)	
}

// ------------------------------------------------------------------
type String16 struct {
	GoValue string 
}


// ------------------------------------------------------------------
type String32 struct {
	GoValue string 
}

func (self String32) SqlType() string {
	return "TEXT"
}

func (self String32) SqlVal() string {
	return quote(self.GoValue)	
}


// ------------------------------------------------------------------
type String64 struct {
	GoValue string 
}

type TinyBlob struct {
	GoValue []byte 
}

type Blob struct {
	GoValue []byte 
}

type MediumBlob struct {
	GoValue []byte 
}

// ------------------------------------------------------------------
type LongBlob struct {
	GoValue []byte
}

func (self LongBlob) SqlType() string { 
	return "LONGBLOB" 
}

func (self LongBlob) SqlVal() string {
	return quote(string(self.GoValue))
}


// ------------------------------------------------------------------
type DateTime struct {
    GoValue time.Time	
}
type Date struct {
	GoValue time.Time 
}
type Time struct {
	GoValue time.Time 
}

type SqlType struct { 
	Val interface{}
}

type Field struct {
	Name string
	Type SqlType 
}

type TableI interface {
	TableName() string
	GetFields() []FieldI
	GetForeignFields() []TableI
}

type FieldI interface {	
	SqlType() string
	SqlVal() string
}


// No Magic
// DRY
// immutable state