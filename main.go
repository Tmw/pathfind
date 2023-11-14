package main

import (
	"fmt"

	"github.com/tmw/pathfind/pkg/prioqueue"
)

type Fruit struct {
	name string
}

func main() {
	// party with guests ranked by popularity?
	pq := prioqueue.New[Fruit]()

	pq.Push(Fruit{"Apple"}, 10)
	pq.Push(Fruit{"pear"}, 6)
	pq.Push(Fruit{"strawberry"}, 8)
	pq.Push(Fruit{"melon"}, 17)

	for pq.Len() > 0 {
		item := pq.PopItem()
		fmt.Printf("%s => %d\n", item.Value.name, item.Priority())
	}
}
