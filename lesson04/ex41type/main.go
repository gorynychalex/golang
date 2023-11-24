package main

import "fmt"

type library []string

func printBooks(lib library){
	for idx, value := range lib {
		fmt.Println((idx+1)," - ", value)
	}
}

func main()  {
	
	var myLibrary library = library{"Tails","Novel","Story"}

	printBooks(myLibrary)

}