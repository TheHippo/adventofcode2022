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

type Group struct {
	rucksacks [3][]byte
}

func (r Group) findCommon() (result int) {
	for _, v := range r.rucksacks[0] {
		if bytes.IndexByte(r.rucksacks[1], v) != -1 && bytes.IndexByte(r.rucksacks[2], v) != -1 {
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

	groups := make([]*Group, 0)

	group := &Group{
		rucksacks: [3][]byte{},
	}

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		size := len(line)
		if size%2 != 0 {
			log.Fatal("line does not contain even number of items")
		}
		c := []byte{}
		for i := 0; i < size; i++ {
			c = Put(c, line[i])
		}
		group.rucksacks[count%3] = c

		count++
		if count%3 == 0 && count != 0 {
			groups = append(groups, group)
			group = &Group{
				rucksacks: [3][]byte{},
			}
		}

	}
	groups = append(groups, group)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result int
	for _, r := range groups {
		result += r.findCommon()
	}
	log.Println(result)

}
