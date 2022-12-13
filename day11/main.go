package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// sorry, but I was too lazy
// https://regex101.com/r/jqBQES/1
var monkey = regexp.MustCompile(`Monkey\s(\d+):\s+Starting items:\s([0-9, ]+)\s+Operation:\snew = old ([*+])\s(old|\d+)\s+Test: divisible by (\d+)\s+If true: throw to monkey (\d+)\s+If false: throw to monkey (\d+)`)

func parseInt(str string) int {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalf("could not parse int: %s - %s", str, err.Error())
	}
	return int(i)
}

func parseList(str string) []int {
	parts := strings.Split(str, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		result = append(result, parseInt(strings.TrimSpace(part)))
	}
	return result
}

type Operation interface {
	apply(old int) int
}

type BaseOperation struct {
	baseValue int
	useOld    bool
}

type Plus struct {
	BaseOperation
}

func (p *Plus) apply(old int) int {
	if p.useOld {
		return old + old
	}
	return old + p.baseValue
}

type Multiply struct {
	BaseOperation
}

func (m *Multiply) apply(old int) int {
	if m.useOld {
		return old * old
	}
	return old * m.baseValue
}

type Monkey struct {
	index       int
	items       []int
	operation   Operation
	test        int
	trueMonkey  int
	falseMonkey int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	matches := monkey.FindAllSubmatch(content, -1)
	for _, match := range matches {
		if len(match) != 8 {
			log.Fatalf("Invalid monkey: %s", string(match[0]))
		}
		var o Operation
		switch string(match[3]) {
		case "+":
			if string(match[4]) == "old" {
				o = &Plus{
					BaseOperation: BaseOperation{
						baseValue: 0,
						useOld:    true,
					},
				}
			} else {
				o = &Plus{
					BaseOperation: BaseOperation{
						baseValue: parseInt(string(match[4])),
						useOld:    false,
					},
				}
			}

		case "*":
			if string(match[4]) == "old" {
				o = &Multiply{
					BaseOperation: BaseOperation{
						baseValue: 0,
						useOld:    true,
					},
				}
			} else {
				o = &Multiply{
					BaseOperation: BaseOperation{
						baseValue: parseInt(string(match[4])),
						useOld:    false,
					},
				}
			}
		default:
			log.Fatal("unknown operation")
		}

		m := &Monkey{
			index:       parseInt(string(match[1])),
			items:       parseList(string(match[2])),
			operation:   o,
			test:        parseInt(string(match[5])),
			trueMonkey:  parseInt(string(match[6])),
			falseMonkey: parseInt(string(match[7])),
		}

		log.Printf("%+v", m)

	}
}
