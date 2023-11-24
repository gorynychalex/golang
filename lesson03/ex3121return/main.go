package main

import "fmt"

func main() {
    a := add(4, 5)   // 9
    b := add(20, 6)  // 26

    fmt.Println(a)
    fmt.Println(b)
}

func add(x, y int) int {
    return x + y
}