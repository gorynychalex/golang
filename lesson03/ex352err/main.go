package main

import (
	"errors"
	"fmt"
)

func main() {

	err := errors.New("my generated error")

	fmt.Println(err)
}
