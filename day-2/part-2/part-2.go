package main

import (
	"fmt"
	"os"
	"bufio"
)

// shape scores
const (
	ROCK int = 1
	PAPER = 2
	SCISSORS = 3
)

// game scores
const (
	LOSS int = 0
	DRAW = 3
	WIN = 6
)

// opponent and player mappings
const (
	A int = ROCK
	B = PAPER
	C = SCISSORS
)

// player mappings
const (
	X int = LOSS
	Y = DRAW
	Z = WIN
)

func main() {
	file, e := os.Open("./input.txt")
	if e != nil {
		panic(e)
	}
	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		opponent := getMove(line[0])
		player := getMove(line[2])
		score += play(player, opponent)
	}
	fmt.Println(score) 
}

func getMove(move byte) int {
	switch move {
	case 'A':
		return A
	case 'B':
		return B
	case 'C':
		return C
	case 'X':
		return X
	case 'Y':
		return Y
	case 'Z':
		return Z
	default:
		panic(string(move) + " is not a valid move")
	}
}

func play(player int, opponent int) int {
	loseLut := []int{SCISSORS, ROCK, PAPER}
	winLut := []int{PAPER, SCISSORS, ROCK}

	if player == LOSS {
		return LOSS + loseLut[opponent - 1]
	}
	if player == WIN{
		return WIN + winLut[opponent - 1]
	}
	return DRAW + opponent
}
