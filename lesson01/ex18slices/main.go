package lesson01

import "fmt"

func main() {
	

	users := [5]string {"Vasya", "Petya", "Kolya", "Olya", "Masha"}

	users1 := users[0:2]
	users2 := users[2:4]
	users3 := users[:3]

	fmt.Println(users,len(users),cap(users))
	fmt.Println(users1,len(users1),cap(users1))
	fmt.Println(users2,len(users2),cap(users2))
	fmt.Println(users3,len(users3),cap(users3))

	users4 := append(users[:],"Maksim","Ivan")
	fmt.Println(users4,len(users4),cap(users4))

	users5 := append(users4[0:2],users4[4:]...)
	fmt.Println(users5,len(users5),cap(users5))

	users10 := make([]string, len(users4),len(users4))
	copy(users10,users4)
	fmt.Println(users10,len(users10),cap(users10))
	// fmt.Println(users10[9])

	for _, value := range users {
		fmt.Println(" - ", value)
	}


}