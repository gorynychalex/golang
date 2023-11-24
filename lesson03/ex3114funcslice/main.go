package main

import "fmt"

func main() {
	nums:=[]int{5,8,34,12}

	sum(nums)
}

func sum(numbers []int){
	var sum = 0
    for _, number := range numbers{
        sum += number
    }
    fmt.Println(sum)
}