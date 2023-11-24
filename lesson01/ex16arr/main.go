package lesson01

import "fmt"

func main()  {
	

	var arr [5]int = [5]int {41,32,53,94,15}

	// fmt.Println(arr[0])
	// fmt.Println(arr[3])

	for idx, n := range arr {
		fmt.Println("arr[",idx,"] = ",n)
	}

}