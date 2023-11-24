package main

import (
	"fmt"
	"math"
)

func main() {
	
	a:=8
	b:=2
	lim := 10000

	if a > b { 
		fmt.Println("а больше б") 
	} else if a < b { 
		fmt.Println("а меньше б") 
	} else { 
		fmt.Println("какой-то другой случай")
	}

	if v := math.Pow(float64(a),float64(b)); v < float64(lim) {
		fmt.Println(v)
	}

}