package main

import (
	"fmt"
	"math"
)

func main() {
var index int8 = 15
var bigIndex int32
bigIndex = int32(index)

// Выведем:
fmt.Println(bigIndex)         // 15
fmt.Printf("%T \n", bigIndex) // int32

// По аналогии выше легко понять как конвертировать в другие типы:
var a int32 = 22
var b uint64
b = uint64(a)

// Выведем
fmt.Println(b)         // 22
fmt.Printf("%T \n", b) // uint64

var big int64 = 264

var little int8

little = int8(big)

fmt.Println(little) //64


fmt.Println(math.MaxInt8)   // 127
fmt.Println(math.MaxUint8)  // 255
fmt.Println(math.MaxInt16)  // 32767
fmt.Println(math.MaxUint16) // 65535
fmt.Println(uint64(math.MaxUint64))

}