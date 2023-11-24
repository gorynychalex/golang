package main

import "fmt"
 
func action(n1 int, n2 int, operation func(int, int) int){ 
    result := operation(n1, n2)
    fmt.Println(result)
}

func main() {

    action(10, 25, func(a,b int) int { return a * b })  // 250
}