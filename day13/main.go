package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
)

var findList = regexp.MustCompile(`\[(.*)\]`)

type IntValue struct {
	value int
}

type ListValue struct {
	value []*IntValue
}

type Value interface {
	CompareTo(Value) bool
}

func (i *IntValue) CompareTo(t Value) bool {
	switch v := t.(type) {
	case *IntValue:
		if v.value <= i.value {
			return true
		} else {
			return false
		}
	case *ListValue:
		l := &ListValue{
			value: []*IntValue{
				i,
			},
		}
		return l.CompareTo(v)
	}
	log.Fatal("not possible")
	return false
}

func (l *ListValue) CompareTo(t Value) bool {
	switch v := t.(type) {
	case *IntValue:
		newL := &ListValue{
			value: []*IntValue{v},
		}
		return l.CompareTo(newL)
	case *ListValue:
		min := int(math.Min(float64(len(l.value)), float64(len(v.value))))
		for i := 0; i < min; i++ {
			if !l.value[i].CompareTo(v.value[i]) {
				return false
			}
		}
		if len(l.value) <= len(v.value) {
			return true
		}
		return false
	}
	log.Fatal("not possible")
	return false
}

type Expression struct {
	Value []Value
}

func parseExpression(str string) *Expression {
	expr := &Expression{
		Value: []Value{},
	}
	return expr
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
		if len(line) > 0 {
			log.Printf("%+v", parseExpression(line))
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
