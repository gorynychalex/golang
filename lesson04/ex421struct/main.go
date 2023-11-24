package main

import (
	"fmt"
	"math"
)

type Point struct {
	x,y float32
}

type Circle struct {
	r float32
	point Point
}

type Rectangle struct {
	a, b float32
	point Point
}

// func area(c *Circle) float32 {
// 	return math.Pi * c.r * c.r
// }


func (c *Circle) area() float32 {
    return math.Pi * c.r * c.r
}

func main()  {


	// var c_new Circle

	// fmt.Println("c_new(): ",c_new)

	// p := new(Point); fmt.Println(p)

	// c := new(Circle); fmt.Println(c)

	p := Point{10,20}


	// rect1 := Rectangle{a:100,b:150}

	circle1 := Circle{
		r : 10,
		point : p,
	}

	circle2 := Circle{100, Point{5,7}}

	

	fmt.Println(circle1)
	// fmt.Println(circle1.r)
	// fmt.Println(circle1.point)
	// fmt.Println(circle1.point.x)
	// fmt.Println(circle1.point.y)
	fmt.Println(circle2)

	// fmt.Println(area(&circle1))
	// fmt.Println(area(&circle2))
	

	area1 := circle1.area()
	area2 := circle1.area()

	fmt.Println(area1)
	fmt.Println(area2)
}