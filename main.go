package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/fareez-ahamed/go-dsa/queue"
)

func main() {
	q := queue.SafeQueue[int]{}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(jobNo int) {
			defer wg.Done()
			addRandomToQueue(jobNo, &q, 5000)
		}(i)
	}
	wg.Wait()
	var i int
	for {
		if val, ok := q.Dequeue(); ok {
			fmt.Printf("%d - %d\n", i, val)
			i++
		} else {
			break
		}
	}
}

func addRandomToQueue(jobNo int, q *queue.SafeQueue[int], count int) {
	fmt.Printf("Job %d Started\n", jobNo)
	defer fmt.Printf("Job %d Completed\n", jobNo)
	for i := 0; i < count; i++ {
		q.Enqueue(rand.Int())
	}
}
