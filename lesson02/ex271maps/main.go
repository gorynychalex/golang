package main

import "fmt"

func main() {
	
	users := make(map[int]string)

	users[100] = "Pupkin"
	users[200] = "Ivanov"
	users[300] = "Sidorov"

	// fmt.Println(users)


	v, ok := users[100]
	fmt.Println("The value:", v, "Present?", ok)

	delete(users,100)

	v, ok = users[100]
	fmt.Println("The value:", v, "Present?", ok)


	for key, value := range users {
		fmt.Println(key, value)
	}
}