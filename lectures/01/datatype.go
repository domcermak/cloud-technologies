package main

import (
	"fmt"
	"math"
)

func main() {

	// bool
	var b = false

	// string
	var s = "Hello"

	// [-128,127], try 258
	var i8 int8 = 127

	var i16 int16 = math.MaxInt16
	var i32 int32 = math.MinInt32
	var i64 int64 = math.MaxInt64

	var u8 uint8 = 255

	var u16 uint16 = math.MaxUint16
	var u32 uint32 = math.MaxUint8
	var u64 uint64 = math.MaxUint64

	// alias for uint8
	var by byte = 127

	var ru rune = 'a'

	var f32 = 3.15
	var f64 = 1e7

	//
	var cmpx = complex(10, 11)

	fmt.Println("bool", b)
	fmt.Println("string", s)
	fmt.Println("int", i8, i16, i32, i64)
	fmt.Println("uint", u8, u16, u32, u64)
	fmt.Println("alias", by, ru)
	fmt.Println("float", f32, f64)
	fmt.Println("complex", cmpx)
}
