package main

import (
	"bufio"
	"log"
	"os"
)

type Play uint8

const (
	_ Play = iota
	Rock
	Paper
	Scissors
)

func (p Play) String() string {
	switch p {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return "unknown"
	}
}

func parsePlay(s rune) Play {
	switch s {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	default:
		log.Fatalf("Unknown play: %v", s)
	}
	return 0
}

func getScore(opponent, you Play) (result int) {
	result = int(you)
	if opponent == you {
		result += 3
		return
	}
	if you == Paper && opponent == Rock {
		result += 6
		return
	}
	if you == Scissors && opponent == Paper {
		result += +6
		return
	}
	if you == Rock && opponent == Scissors {
		result += 6
		return
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	total := 0
	for reader.Scan() {
		line := []rune(reader.Text())
		if len(line) != 3 {
			log.Fatal("Wrong line lenghth")
		}
		opponentPlay, yourPlay := parsePlay(line[0]), parsePlay(line[2])
		total += getScore(opponentPlay, yourPlay)
	}
	log.Println(total)

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
}
