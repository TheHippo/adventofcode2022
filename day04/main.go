package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var findParts = regexp.MustCompile(`(\d{1,2})-(\d{1,2}),(\d{1,2})-(\d{1,2})`)

// .2345678.  2-8
// ..34567..  3-7

// .....6...  6-6
// ...456...  4-6

func parse(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}

func within(e1s, e1e, e2s, e2e int) bool {
	if e1e >= e2s && e1s <= e2e {
		return true
	}
	if e2e >= e1s && e2s <= e1e {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := findParts.FindAllStringSubmatch(line, -1)
		if len(parts) != 1 && len(parts[0]) != 5 {
			log.Fatal("Invalid line")
		}
		e1s, e1e, e2s, e2e := parse(parts[0][1]), parse(parts[0][2]), parse(parts[0][3]), parse(parts[0][4])
		if within(e1s, e1e, e2s, e2e) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(count)
}
