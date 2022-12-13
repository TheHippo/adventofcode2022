package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	X       int
	Cycle   int
	History []int
}

func (m *Machine) Scan(line string) {
	if strings.HasPrefix(line, "noop") {
		m.Cycle++
		m.History = append(m.History, m.X)
	} else if strings.HasPrefix(line, "addx") {
		i, err := strconv.ParseInt(line[5:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		m.Cycle += 2
		m.History = append(m.History, m.X, m.X)
		m.X += int(i)
	} else {
		log.Fatalf("Invalid instruction: %s", line)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := &Machine{
		X:       1,
		Cycle:   0,
		History: []int{},
	}

	for scanner.Scan() {
		line := scanner.Text()
		m.Scan(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lines := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, i := range lines {
		s := (i) * m.History[i-1]
		sum += s
		log.Printf("%3d: %d", i, s)
	}
	log.Println(sum)

	str := ""
	for i := 0; i < len(m.History); i++ {
		var spriteStart, spriteEnd = (i%40 - 1), (i%40 + 1)
		if spriteStart <= m.History[i] && m.History[i] <= spriteEnd {
			str += "#"
		} else {
			str += "."
		}
		if (i+1)%40 == 0 {
			log.Println(str)
			str = ""
		}
	}

}
