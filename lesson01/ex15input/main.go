package lesson01

import "fmt"

func main()  {
	var name string
	var age int

	fmt.Print("Введите имя:")
	fmt.Scan(&name)
	fmt.Print("Введите возраст:")
	fmt.Scan(&age)

	fmt.Println("Привет ", name)
	fmt.Printf("Тебе %f", float32(age))

}