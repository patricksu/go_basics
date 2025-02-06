package main

import (
	"fmt"

	"github.com/patricksu/go_basics/ds"
)

func main() {
	fmt.Println("Hello, World!")
	stack := ds.Stack{
		Items: []int{1, 2, 3},
	}
	fmt.Println(stack)
}
