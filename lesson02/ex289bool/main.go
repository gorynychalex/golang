package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "true"
	res, err := strconv.ParseBool(s)
	
	if err != nil { // не забываем проверить ошибку
		panic(err)
	}
	fmt.Println(res)      // true
	fmt.Printf("%T", res)  // bool
}