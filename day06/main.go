package main

import (
	"bufio"
	"log"
	"os"
)

type Ring struct {
	content []byte
	size    int
}

// x 1 2 3 4
// 1 x
// 2 x x
// 3 x x x
// 4 x x x x

func (r Ring) allDifferent() bool {
	for i := 0; i < r.size-1; i++ {
		for j := i + 1; j < r.size; j++ {
			if r.content[i] == r.content[j] {
				return false
			}
		}
	}
	return true
}

func newRing(size int, str string) *Ring {
	if len(str) < size {
		log.Fatal("line too short")
	}
	return &Ring{
		size:    size,
		content: []byte(str[0:size]),
	}
}

func (r *Ring) push(b byte) {
	r.content = append(r.content[1:], b)
}

func (r *Ring) String() string {
	return string(r.content)
}

func getStart(line string) (result int) {
	ring := newRing(4, line)
	result = 4
	for {
		if ring.allDifferent() {
			return
		}
		result++
		if result > len(line) {
			log.Fatal("invalid input")
		}
		ring.push(line[result-1])
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("%d - %s", getStart(line), line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
