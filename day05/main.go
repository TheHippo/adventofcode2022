package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack []byte

type State struct {
	stacks []Stack
}

var findMove = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := &State{
		stacks: []Stack{},
	}

	stackCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(strings.TrimLeft(line, " "), "[") {
			// end of state
			break
		}
		if len(state.stacks) == 0 {
			stackCount = (len(line) + 1) / 4
			state.stacks = make([]Stack, stackCount)
		}
		for i := 0; i < stackCount; i++ {
			part := line[i*4 : i*4+3]
			if strings.HasPrefix(part, "[") && strings.HasSuffix(part, "]") {
				state.stacks[i] = append([]byte{part[1]}, state.stacks[i]...)
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "m") {
			continue
		}
		matches := findMove.FindStringSubmatch(line)
		if len(matches) != 4 {
			log.Fatal("invalid move line")
		}
		count, from, to := parseInt(matches[1]), parseInt(matches[2])-1, parseInt(matches[3])-1
		if len(state.stacks[from]) < count {
			log.Fatalf("cannot move that much items: %s", line)
		}
		var moving []byte
		moving, state.stacks[from] = state.stacks[from][len(state.stacks[from])-count:], state.stacks[from][:len(state.stacks[from])-count]
		for left, right := 0, len(moving)-1; left < right; left, right = left+1, right-1 {
			moving[left], moving[right] = moving[right], moving[left]
		}
		state.stacks[to] = append(state.stacks[to], moving...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := strings.Builder{}
	for _, s := range state.stacks {
		result.WriteByte(s[len(s)-1])
	}
	log.Println(result.String())

}
