package main

import (
	"fmt"

	"github.com/fareez-ahamed/go-dsa/queue"
)

func main() {
	q := queue.Queue[int]{}
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	for j := 0; j < 10; j++ {
		if val, ok := q.Dequeue(); ok {
			fmt.Println(val)
		} else {
			break
		}
	}
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	for j := 0; j < 10; j++ {
		if val, ok := q.Dequeue(); ok {
			fmt.Println(val)
		} else {
			break
		}
	}
}
