package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm"
)

func main() {
	s := dataStructureAlgorithm.NewStack()
	s.Push("china")
	s.Push("usa")

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
