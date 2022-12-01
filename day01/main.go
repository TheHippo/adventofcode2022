package main

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := bufio.NewScanner(file)
	current := 0

	elves := &IntHeap{}
	heap.Init(elves)

	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			heap.Push(elves, current)
			current = 0
		} else {
			parsed, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			current += int(parsed)
		}
	}

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
	max := 0

	for i := 0; i < 3; i++ {
		// log.Println(heap.Pop(elves))
		max += heap.Pop(elves).(int)
	}

	log.Println(max)

}
