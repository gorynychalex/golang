package lesson01

import "fmt"

func main()  {
	
	a:=16

	b:=20

	c:=a+b

	fmt.Println(c)

	if a%2==0 {
		fmt.Println("a - четное число")
	} else {
		fmt.Println("a - НЕ четное число")
	}

	fmt.Println("последний разряд числа а = ",a%10)


}