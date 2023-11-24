package main

import (
	"fmt"
	"reflect"
)

type MyInt int

func main() {
	

var x MyInt = 12
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type().Name())                   // MyInt.
fmt.Println("kind is uint8: ", v.Kind() == reflect.Int) // true.
}