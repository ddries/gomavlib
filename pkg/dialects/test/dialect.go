//nolint:golint,misspell,govet// Package test contains the test dialect (autogenerated).
package test

import (
	"github.com/aler9/gomavlib/pkg/dialect"
	"github.com/aler9/gomavlib/pkg/msg"
)

// Dialect contains the dialect object that can be passed to the library.
var Dialect = dial

// dialect is not exposed directly such that it is not displayed in godoc.
var dial = &dialect.Dialect{3, []msg.Message{
	// test.xml
	&MessageTestTypes{},
}}

// test.xml

// Test all field types
type MessageTestTypes struct {
	// char
	C string
	// string
	S string `mavlen:"10"`
	// uint8_t
	U8 uint8
	// uint16_t
	U16 uint16
	// uint32_t
	U32 uint32
	// uint64_t
	U64 uint64
	// int8_t
	S8 int8
	// int16_t
	S16 int16
	// int32_t
	S32 int32
	// int64_t
	S64 int64
	// float
	F float32
	// double
	D float64
	// uint8_t_array
	U8Array [3]uint8
	// uint16_t_array
	U16Array [3]uint16
	// uint32_t_array
	U32Array [3]uint32
	// uint64_t_array
	U64Array [3]uint64
	// int8_t_array
	S8Array [3]int8
	// int16_t_array
	S16Array [3]int16
	// int32_t_array
	S32Array [3]int32
	// int64_t_array
	S64Array [3]int64
	// float_array
	FArray [3]float32
	// double_array
	DArray [3]float64
}

// GetID implements the msg.Message interface.
func (*MessageTestTypes) GetID() uint32 {
	return 17000
}
