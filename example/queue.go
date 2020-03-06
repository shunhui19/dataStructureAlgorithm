package main

import (
	"fmt"

	"github.com/shunhui19/dataStructureAlgorithm"
)

func main() {
	q := dataStructureAlgorithm.NewQueue(5)
	q.Enqueue("china")
	q.Enqueue("usa")
	q.Enqueue("england")
	q.Enqueue("india")
	fmt.Println(q.Dequeue())
	q.Enqueue("frank")

	q.MakeEmpty()
	fmt.Println(q.Dequeue())
}
