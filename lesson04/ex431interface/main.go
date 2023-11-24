package main

import "fmt"

type Vehicle interface{
    move()
}

func drive(v Vehicle) { v.move() }

// структура "Автомобиль"
type Car struct{ }
 
// структура "Самолет"
type Aircraft struct{}
 
 
func (c Car) move(){
    fmt.Println("Автомобиль едет")
}
func (a Aircraft) move(){
    fmt.Println("Самолет летит")
}
 
func main() {
     
    var tesla Vehicle = Car{}
    var boing Vehicle = Aircraft{}
    // drive(tesla)
    // drive(boing)

    vs := []Vehicle{tesla,boing}

    for _,v := range vs {
        drive(v)
    }
}