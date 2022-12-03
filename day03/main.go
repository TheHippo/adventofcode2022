package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func Value(i byte) int {
	if i >= 65 && i <= 90 {
		// upper case letter
		return int(i - 38)
	}
	if i >= 97 && i <= 122 {
		// lower case latter
		return int(i - 96)
	}
	panic("out of range item")
}

func Put(c []byte, i byte) []byte {
	if bytes.IndexByte(c, i) == -1 {
		c = append(c, i)
	}
	return c
}

type Rucksack struct {
	compartments [2][]byte
}

func (r Rucksack) findCommon() (result int) {
	for _, v := range r.compartments[0] {
		if bytes.IndexByte(r.compartments[1], v) != -1 {
			result = +Value(v)
		}
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rucksacks := make([]*Rucksack, 0)

	for scanner.Scan() {
		line := scanner.Text()
		size := len(line)
		if size%2 != 0 {
			log.Fatal("line does not contain even number of items")
		}
		c := []byte{}
		rucksack := &Rucksack{
			compartments: [2][]byte{nil, nil},
		}

		for i := 0; i < size; i++ {
			if i == 0 || i == (size/2) {
				if i != 0 {
					rucksack.compartments[0] = c
				}
				c = []byte{}
			}
			c = Put(c, line[i])
		}
		rucksack.compartments[1] = c
		rucksacks = append(rucksacks, rucksack)
	}

	var result int
	for _, r := range rucksacks {
		result += r.findCommon()
	}
	log.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
