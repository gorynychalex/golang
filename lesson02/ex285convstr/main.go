package main

import (
	"fmt"
	"strconv"
)

func main() {
	user := "Пупкин"
	code := 115

	fmt.Printf("%T \n", code)
	fmt.Printf("%T \n", strconv.FormatInt(int64(code),10))

	result := "Ученик " + user + " с кодом " + strconv.Itoa(code)

	fmt.Println(result)


	i := 0b01010101
	fmt.Println(i)
	fmt.Println(strconv.FormatInt(int64(i), 2))

	var a int64 = 0xB // 'B' в шестнадцатеричной это 11 в десятичной системе
	fmt.Println(a)
	fmt.Println(strconv.FormatInt(a, 10)) // 11
	fmt.Println(strconv.FormatInt(a, 16)) // b
	
	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).

	var f float64 = 1.0123456789
	fmt.Println(strconv.FormatFloat(f, 'e', -1, 64))

	b := "true"
	res, err := strconv.ParseBool(b)
	if err != nil { panic(err) }
	fmt.Println(res)

	


}