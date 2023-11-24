package main

import "fmt"

func main() {
    add("sum =",1, 2, 3)        // sum = 6
    add("sum = ",1, 2, 3, 4)     // sum = 10
    add("mysum =",5, 6, 7, 2, 3)  // sum = 23

	nums:=[]int{5,8,34,12}
	add("sum = ",nums...)
}
 
func add(s string, numbers ...int){
    var sum = 0
    for _, number := range numbers{
        sum += number
    }
    fmt.Println(s, sum)
}

