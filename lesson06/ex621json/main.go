package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"firstname"`
	Age    int
	Status bool
	Values []int
}

func main() {

	var x interface{} 
	x = "mystring" // в памяти ("mystring", string)

	s := Person {
		Name: "Alex",
		Age: 25,
		Status: true,
		Values: []int{1,2,3},
	}

	data, err := json.Marshal(s)
	if err != nil { fmt.Println(err); return }
	fmt.Printf("%s\n",&data)

	var s1 Person

	json.Unmarshal(data,&s1)

	fmt.Printf("%v",s1)
}