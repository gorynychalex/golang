package main

import "fmt"

func add(x int, y int) int{ return x + y}
func subtract(x int, y int) int{ return x - y}
func multiply(x int, y int) int{ return x * y}

func selectFunc(n int) (func(int, int) int){
    if n==1 {
        return add
    }else if n==2{
        return subtract
    }else{
        return multiply
    }
}

func main() {

	
	fmt.Println(selectFunc(0)(5,10))
	

}
