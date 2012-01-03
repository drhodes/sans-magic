package sansmagic

import (
	//"time"
	//"fmt"
)

// // Numeric Types ------------------------------------------------------------------
// //type Int8 struct {
// //	GoVal int8
// //}
// type Int8 struct {
// 	GoVal int8
// }

// func (self Int8) Copy() FieldI {
// 	return Int8{self.GoVal}
// }

// func (self Int8) SqlType() string {
// 	return "TINYINT"
// }

// func (self Int8) SqlVal() string {
// 	return quote(fmt.Sprintf("%d", self.GoVal))
// }

// // -----------------------------------------------------------------
// type Int16 struct {
// 	GoVal int16
// }
// type Int32 struct {
// 	GoVal int32
// }
// type Int64 struct {
// 	GoVal int64
// }
// type Uint8 struct {
// 	GoVal uint8
// }
// type Uint16 struct {
// 	GoVal uint16
// }
// type Uint32 struct {
// 	GoVal uint32
// }
// type Uint64 struct {
// 	GoVal uint64
// }
// type Float32 struct {
// 	GoVal float32
// }
// type Float64 struct {
// 	GoVal float64
// }
// type Bool struct {
// 	GoVal bool
// }
// /*
// type String struct {
// 	GoVal string 
// }
// */
// // -----------------------------------------------------------------
// type String8 struct {
// 	GoVal string
// }

// func (self String8) Copy() FieldI {
// 	return String8{self.GoVal}
// }
// func (self String8) SqlType() string {
// 	return "TINYTEXT"
// }
// func (self String8) SqlVal() string {
// 	return quote(string(self.GoVal))
// }
// // ------------------------------------------------------------------
// type String16 struct {
// 	GoVal string
// }
// // ------------------------------------------------------------------
// type String32 struct {
// 	GoVal string
// }

// func (self String32) Copy() FieldI {
// 	return String32{self.GoVal}
// }
// func (self String32) SqlType() string {
// 	return "TEXT"
// }
// func (self String32) SqlVal() string {
// 	return quote(self.GoVal)
// }
// // ------------------------------------------------------------------
// type String64 struct {
// 	GoVal string
// }
// type TinyBlob struct {
// 	GoVal []byte
// }
// type Blob struct {
// 	GoVal []byte
// }
// type MediumBlob struct {
// 	GoVal []byte
// }
// // -----------------------------------------------------------------
// type LongBlob []byte

// func (self LongBlob) Copy() FieldI {
// 	return LongBlob(self)
// }
// func (self LongBlob) SqlType() string {
// 	return "LONGBLOB"
// }
// func (self LongBlob) SqlVal() string {
// 	return quote(fmt.Sprintf("%v", self))
// }
// // ------------------------------------------------------------------
// type DateTime struct {
// 	GoVal time.Time
// }

// type Date struct {
// 	GoVal time.Time
// }
// type Time struct {
// 	GoVal time.Time
// }

// type SqlType struct {
// 	Val interface{}
// }

// No Magic
// DRY
// immutable state
