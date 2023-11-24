package main

import "fmt"


const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thirthday
	Friday
	Saturday
)

func main(){

	const one = 30

	const abc int = one

	const (
		Pi float32 = 3.1415
		e = 2.71
	)

	var a int = 25

	fmt.Println(abc)
	fmt.Println(a)
	fmt.Println(Pi)
	fmt.Println(e)
	
	fmt.Println("Sunday is ",Sunday, " day of week")

}