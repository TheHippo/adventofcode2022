package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := bufio.NewScanner(file)
	current, max := 0, 0

	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			if current > max {
				max = current
			}
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
	log.Println(max)
}
