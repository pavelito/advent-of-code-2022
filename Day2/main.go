package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	pointMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
		"A": 1,
		"B": 2,
		"C": 3,
	}

	winnerMap := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	drawMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	playMap := map[string]map[string]string{
		"A": {"X": "C", "Y": "A", "Z": "B"},
		"B": {"X": "A", "Y": "B", "Z": "C"},
		"C": {"X": "B", "Y": "C", "Z": "A"},
	}

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	strategyScorePart1 := 0
	strategyScorePart2 := 0

	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		moves := strings.Fields(stringVal)
		roundScorePart1 := 0
		roundScorePart2 := 0

		//Part 1
		if moves[1] == drawMap[moves[0]] {
			fmt.Println("Draw")
			roundScorePart1 += 3
		} else if moves[1] == winnerMap[moves[0]] {
			fmt.Println("Win")
			roundScorePart1 += 6
		} else {
			fmt.Println("Lost")
		}
		roundScorePart1 += pointMap[moves[1]]
		strategyScorePart1 += roundScorePart1

		//Part 2
		moveToPlay := playMap[moves[0]][moves[1]]
		if moves[1] == "X" {
			//loose
			roundScorePart2 += 0
		} else if moves[1] == "Y" {
			//draw
			roundScorePart2 += 3
		} else {
			//win
			roundScorePart2 += 6
		}
		roundScorePart2 += pointMap[moveToPlay]

		strategyScorePart2 += roundScorePart2

		fmt.Println("Moves", moves)
		fmt.Println("Move to play Part 2", moveToPlay)
		fmt.Println("Round Score Part 1", roundScorePart1)
		fmt.Println("Round Score Part 2", roundScorePart2)
	}

	fmt.Println("Strategy Score Part 1", strategyScorePart1)
	fmt.Println("Strategy Score Part 2", strategyScorePart2)

	readFile.Close()
}
