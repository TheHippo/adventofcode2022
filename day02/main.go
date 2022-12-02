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

type Outcome uint8

const (
	_ Outcome = iota
	Lose
	Draw
	Win
)

func parseOutcome(s rune) Outcome {
	switch s {
	case 'X':
		return Lose
	case 'Y':
		return Draw
	case 'Z':
		return Win
	default:
		log.Fatalf("Unknown play: %v", s)
	}
	return 0
}

func (p Outcome) String() string {
	switch p {
	case Lose:
		return "Lose"
	case Draw:
		return "Draw"
	case Win:
		return "Win"
	default:
		return "unknown"
	}
}

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
	case 'A':
		return Rock
	case 'B':
		return Paper
	case 'C':
		return Scissors
	default:
		log.Fatalf("Unknown play: %v", s)
	}
	return 0
}

func choose(opponent Play, outcome Outcome) (result Play) {
	switch outcome {
	case Draw:
		result = opponent
	case Win:
		result = opponent + 1
		if result == 4 {
			result = 1
		}
	case Lose:
		result = opponent - 1
		if result == 0 {
			result = 3
		}
	}
	return
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
		opponentPlay, result := parsePlay(line[0]), parseOutcome(line[2])
		yourPlay := choose(opponentPlay, result)
		total += getScore(opponentPlay, yourPlay)
	}
	log.Println(total)

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
}
