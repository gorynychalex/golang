package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	

	s:= "Привет мой друг!"

	// s[3] = 12 // Ошибка

	fmt.Println("len(): ", len(s))
	fmt.Println("count symbols: ", utf8.RuneCountInString(s))

	println(s)

	for _,c := range s { fmt.Printf("%v ",c)}

	println(len(s))

	println(s[:6])

	s = s[:6] + "вет"

	fmt.Printf("%v\n",s)

	for _,c := range s { fmt.Printf("%v ",c)}

}