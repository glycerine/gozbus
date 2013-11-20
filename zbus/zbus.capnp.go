package gozbus

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/jmckaskill/go-capnproto"
	"math"
	"unsafe"
)

type Zdate C.Struct

func NewZdate(s *C.Segment) Zdate      { return Zdate(s.NewStruct(8, 0)) }
func NewRootZdate(s *C.Segment) Zdate  { return Zdate(s.NewRootStruct(8, 0)) }
func ReadRootZdate(s *C.Segment) Zdate { return Zdate(s.Root(0).ToStruct()) }
func (s Zdate) Year() int16            { return int16(C.Struct(s).Get16(0)) }
func (s Zdate) SetYear(v int16)        { C.Struct(s).Set16(0, uint16(v)) }
func (s Zdate) Month() uint8           { return C.Struct(s).Get8(2) }
func (s Zdate) SetMonth(v uint8)       { C.Struct(s).Set8(2, v) }
func (s Zdate) Day() uint8             { return C.Struct(s).Get8(3) }
func (s Zdate) SetDay(v uint8)         { C.Struct(s).Set8(3, v) }
func (s Zdate) Tjday() uint32          { return C.Struct(s).Get32(4) }
func (s Zdate) SetTjday(v uint32)      { C.Struct(s).Set32(4, v) }

type Zdate_List C.PointerList

func NewZdateList(s *C.Segment, sz int) Zdate_List { return Zdate_List(s.NewUInt64List(sz)) }
func (s Zdate_List) Len() int                      { return C.PointerList(s).Len() }
func (s Zdate_List) At(i int) Zdate                { return Zdate(C.PointerList(s).At(i).ToStruct()) }
func (s Zdate_List) ToArray() []Zdate              { return *(*[]Zdate)(unsafe.Pointer(C.PointerList(s).ToArray())) }

type Ztm C.Struct

func NewZtm(s *C.Segment) Ztm                { return Ztm(s.NewStruct(8, 0)) }
func NewRootZtm(s *C.Segment) Ztm            { return Ztm(s.NewRootStruct(8, 0)) }
func ReadRootZtm(s *C.Segment) Ztm           { return Ztm(s.Root(0).ToStruct()) }
func (s Ztm) TmMsecMidnt() uint32            { return C.Struct(s).Get32(0) }
func (s Ztm) SetTmMsecMidnt(v uint32)        { C.Struct(s).Set32(0, v) }
func (s Ztm) TmNanoSinceLastSec() uint32     { return C.Struct(s).Get32(4) }
func (s Ztm) SetTmNanoSinceLastSec(v uint32) { C.Struct(s).Set32(4, v) }

type Ztm_List C.PointerList

func NewZtmList(s *C.Segment, sz int) Ztm_List { return Ztm_List(s.NewUInt64List(sz)) }
func (s Ztm_List) Len() int                    { return C.PointerList(s).Len() }
func (s Ztm_List) At(i int) Ztm                { return Ztm(C.PointerList(s).At(i).ToStruct()) }
func (s Ztm_List) ToArray() []Ztm              { return *(*[]Ztm)(unsafe.Pointer(C.PointerList(s).ToArray())) }

type Ztd C.Struct

func NewZtd(s *C.Segment) Ztd      { return Ztd(s.NewStruct(0, 2)) }
func NewRootZtd(s *C.Segment) Ztd  { return Ztd(s.NewRootStruct(0, 2)) }
func ReadRootZtd(s *C.Segment) Ztd { return Ztd(s.Root(0).ToStruct()) }
func (s Ztd) Date() Zdate          { return Zdate(C.Struct(s).GetObject(0).ToStruct()) }
func (s Ztd) SetDate(v Zdate)      { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Ztd) Time() Ztm            { return Ztm(C.Struct(s).GetObject(1).ToStruct()) }
func (s Ztd) SetTime(v Ztm)        { C.Struct(s).SetObject(1, C.Object(v)) }

type Ztd_List C.PointerList

func NewZtdList(s *C.Segment, sz int) Ztd_List { return Ztd_List(s.NewCompositeList(0, 2, sz)) }
func (s Ztd_List) Len() int                    { return C.PointerList(s).Len() }
func (s Ztd_List) At(i int) Ztd                { return Ztd(C.PointerList(s).At(i).ToStruct()) }
func (s Ztd_List) ToArray() []Ztd              { return *(*[]Ztd)(unsafe.Pointer(C.PointerList(s).ToArray())) }

type Z C.Struct
type Z_which uint16

const (
	Z_VOID    Z_which = 0
	Z_ZZ              = 1
	Z_F64             = 2
	Z_F32             = 3
	Z_I64             = 4
	Z_I32             = 5
	Z_I16             = 6
	Z_I8              = 7
	Z_U64             = 8
	Z_U32             = 9
	Z_U16             = 10
	Z_U8              = 11
	Z_BOOL            = 12
	Z_TEXT            = 13
	Z_BLOB            = 14
	Z_F64VEC          = 15
	Z_F32VEC          = 16
	Z_I64VEC          = 17
	Z_I32VEC          = 18
	Z_I16VEC          = 19
	Z_I8VEC           = 20
	Z_U64VEC          = 21
	Z_U32VEC          = 22
	Z_U16VEC          = 23
	Z_U8VEC           = 24
	Z_ZVEC            = 25
	Z_ZVECVEC         = 26
	Z_ZDATE           = 27
	Z_ZTM             = 28
	Z_ZTD             = 29
)

func NewZ(s *C.Segment) Z             { return Z(s.NewStruct(16, 1)) }
func NewRootZ(s *C.Segment) Z         { return Z(s.NewRootStruct(16, 1)) }
func ReadRootZ(s *C.Segment) Z        { return Z(s.Root(0).ToStruct()) }
func (s Z) Which() Z_which            { return Z_which(C.Struct(s).Get16(0)) }
func (s Z) Zz() Z                     { return Z(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetZz(v Z)                 { C.Struct(s).Set16(0, 1); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) F64() float64              { return math.Float64frombits(C.Struct(s).Get64(8)) }
func (s Z) SetF64(v float64)          { C.Struct(s).Set16(0, 2); C.Struct(s).Set64(8, math.Float64bits(v)) }
func (s Z) F32() float32              { return math.Float32frombits(C.Struct(s).Get32(8)) }
func (s Z) SetF32(v float32)          { C.Struct(s).Set16(0, 3); C.Struct(s).Set32(8, math.Float32bits(v)) }
func (s Z) I64() int64                { return int64(C.Struct(s).Get64(8)) }
func (s Z) SetI64(v int64)            { C.Struct(s).Set16(0, 4); C.Struct(s).Set64(8, uint64(v)) }
func (s Z) I32() int32                { return int32(C.Struct(s).Get32(8)) }
func (s Z) SetI32(v int32)            { C.Struct(s).Set16(0, 5); C.Struct(s).Set32(8, uint32(v)) }
func (s Z) I16() int16                { return int16(C.Struct(s).Get16(8)) }
func (s Z) SetI16(v int16)            { C.Struct(s).Set16(0, 6); C.Struct(s).Set16(8, uint16(v)) }
func (s Z) I8() int8                  { return int8(C.Struct(s).Get8(8)) }
func (s Z) SetI8(v int8)              { C.Struct(s).Set16(0, 7); C.Struct(s).Set8(8, uint8(v)) }
func (s Z) U64() uint64               { return C.Struct(s).Get64(8) }
func (s Z) SetU64(v uint64)           { C.Struct(s).Set16(0, 8); C.Struct(s).Set64(8, v) }
func (s Z) U32() uint32               { return C.Struct(s).Get32(8) }
func (s Z) SetU32(v uint32)           { C.Struct(s).Set16(0, 9); C.Struct(s).Set32(8, v) }
func (s Z) U16() uint16               { return C.Struct(s).Get16(8) }
func (s Z) SetU16(v uint16)           { C.Struct(s).Set16(0, 10); C.Struct(s).Set16(8, v) }
func (s Z) U8() uint8                 { return C.Struct(s).Get8(8) }
func (s Z) SetU8(v uint8)             { C.Struct(s).Set16(0, 11); C.Struct(s).Set8(8, v) }
func (s Z) Bool() bool                { return C.Struct(s).Get1(64) }
func (s Z) SetBool(v bool)            { C.Struct(s).Set16(0, 12); C.Struct(s).Set1(64, v) }
func (s Z) Text() string              { return C.Struct(s).GetObject(0).ToText() }
func (s Z) SetText(v string)          { C.Struct(s).Set16(0, 13); C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Z) Blob() []byte              { return C.Struct(s).GetObject(0).ToData() }
func (s Z) SetBlob(v []byte)          { C.Struct(s).Set16(0, 14); C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s Z) F64vec() C.Float64List     { return C.Float64List(C.Struct(s).GetObject(0)) }
func (s Z) SetF64vec(v C.Float64List) { C.Struct(s).Set16(0, 15); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) F32vec() C.Float32List     { return C.Float32List(C.Struct(s).GetObject(0)) }
func (s Z) SetF32vec(v C.Float32List) { C.Struct(s).Set16(0, 16); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) I64vec() C.Int64List       { return C.Int64List(C.Struct(s).GetObject(0)) }
func (s Z) SetI64vec(v C.Int64List)   { C.Struct(s).Set16(0, 17); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) I32vec() C.Int32List       { return C.Int32List(C.Struct(s).GetObject(0)) }
func (s Z) SetI32vec(v C.Int32List)   { C.Struct(s).Set16(0, 18); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) I16vec() C.Int16List       { return C.Int16List(C.Struct(s).GetObject(0)) }
func (s Z) SetI16vec(v C.Int16List)   { C.Struct(s).Set16(0, 19); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) I8vec() C.Int8List         { return C.Int8List(C.Struct(s).GetObject(0)) }
func (s Z) SetI8vec(v C.Int8List)     { C.Struct(s).Set16(0, 20); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) U64vec() C.UInt64List      { return C.UInt64List(C.Struct(s).GetObject(0)) }
func (s Z) SetU64vec(v C.UInt64List)  { C.Struct(s).Set16(0, 21); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) U32vec() C.UInt32List      { return C.UInt32List(C.Struct(s).GetObject(0)) }
func (s Z) SetU32vec(v C.UInt32List)  { C.Struct(s).Set16(0, 22); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) U16vec() C.UInt16List      { return C.UInt16List(C.Struct(s).GetObject(0)) }
func (s Z) SetU16vec(v C.UInt16List)  { C.Struct(s).Set16(0, 23); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) U8vec() C.UInt8List        { return C.UInt8List(C.Struct(s).GetObject(0)) }
func (s Z) SetU8vec(v C.UInt8List)    { C.Struct(s).Set16(0, 24); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Zvec() Z_List              { return Z_List(C.Struct(s).GetObject(0)) }
func (s Z) SetZvec(v Z_List)          { C.Struct(s).Set16(0, 25); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Zvecvec() C.PointerList    { return C.PointerList(C.Struct(s).GetObject(0)) }
func (s Z) SetZvecvec(v C.PointerList) {
	C.Struct(s).Set16(0, 26)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Z) Zdate() Zdate     { return Zdate(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetZdate(v Zdate) { C.Struct(s).Set16(0, 27); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Ztm() Ztm         { return Ztm(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetZtm(v Ztm)     { C.Struct(s).Set16(0, 28); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Ztd() Ztd         { return Ztd(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetZtd(v Ztd)     { C.Struct(s).Set16(0, 29); C.Struct(s).SetObject(0, C.Object(v)) }

type Z_List C.PointerList

func NewZList(s *C.Segment, sz int) Z_List { return Z_List(s.NewCompositeList(16, 1, sz)) }
func (s Z_List) Len() int                  { return C.PointerList(s).Len() }
func (s Z_List) At(i int) Z                { return Z(C.PointerList(s).At(i).ToStruct()) }
func (s Z_List) ToArray() []Z              { return *(*[]Z)(unsafe.Pointer(C.PointerList(s).ToArray())) }
