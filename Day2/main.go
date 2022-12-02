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

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	strategyScore := 0
	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		moves := strings.Fields(stringVal)
		roundScore := 0
		if moves[1] == drawMap[moves[0]] {
			fmt.Println("Draw")
			roundScore += 3
		} else if moves[1] == winnerMap[moves[0]] {
			fmt.Println("Win")
			roundScore += 6
		} else {
			fmt.Println("Lost")
		}
		roundScore += pointMap[moves[1]]
		strategyScore += roundScore
		fmt.Println(moves)
		fmt.Println("Round Score", roundScore)
	}

	fmt.Println("Strategy Score", strategyScore)

	readFile.Close()
}
