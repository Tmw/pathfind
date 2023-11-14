package main

import (
	"fmt"

	"github.com/tmw/pathfind/pkg/prioqueue"
)

func main() {
	pq := prioqueue.New()
	pq.Push("banana", 10)

	pq.Push("pear", 6)
	pq.Push("strawberry", 8)
	pq.Push("melon", 17)

	for pq.Len() > 0 {
		item := pq.PopItem()
		fmt.Printf("%s => %d\n", item.Value, item.Priority())
	}
}
