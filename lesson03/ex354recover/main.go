package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			// The return value of `recover` is the error raised in
			// the call to `panic`.
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	
	fmt.Println("After mayPanic()")

}
