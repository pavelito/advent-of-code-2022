package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var tailPositions [][2]int
var headPositions [][2]int
var currentHeadPosition [2]int
var currentTailPosition [2]int

func getUniquePositions(positions [][2]int) [][2]int {
	keys := make(map[string]bool)
	var returnSlice [][2]int
	for _, position := range positions {
		key := fmt.Sprintf("%d-%d", position[0], position[1])
		if _, value := keys[key]; !value {
			keys[key] = true
			returnSlice = append(returnSlice, position)
		}
	}
	return returnSlice
}

func shouldTailMove(headPosition [2]int, tailPosition [2]int) bool {
	if math.Abs(float64(headPosition[0]-tailPosition[0])) == 2 ||
		math.Abs(float64(headPosition[1]-tailPosition[1])) == 2 {
		return true
	}
	return false
}

func move(direction string, steps int) {
	fmt.Println("Moving ", direction, steps)
	for step := 1; step <= steps; step++ {
		//move head
		switch direction {
		case "R":
			currentHeadPosition[1] = currentHeadPosition[1] + 1
		case "U":
			currentHeadPosition[0] = currentHeadPosition[0] + 1
		case "L":
			currentHeadPosition[1] = currentHeadPosition[1] - 1
		case "D":
			currentHeadPosition[0] = currentHeadPosition[0] - 1
		}
		headPositions = append(headPositions, currentHeadPosition)
		//move tail now
		if shouldTailMove(currentHeadPosition, currentTailPosition) {
			//horizontal move
			if currentHeadPosition[0] == currentTailPosition[0] {
				if currentTailPosition[1] < currentHeadPosition[1] {
					currentTailPosition[1]++
				} else {
					currentTailPosition[1]--
				}
			}
			//vertical move
			if currentHeadPosition[1] == currentTailPosition[1] {
				if currentTailPosition[0] < currentHeadPosition[0] {
					currentTailPosition[1]++
				} else {
					currentTailPosition[1]--
				}
			}
			//diagonal moves
			if currentHeadPosition[0]-currentTailPosition[0] == 2 {
				currentTailPosition[0]++
				if currentHeadPosition[1] > currentTailPosition[1] {
					currentTailPosition[1]++
				} else {
					currentTailPosition[1]--
				}
			} else if currentHeadPosition[0]-currentTailPosition[0] == -2 {
				currentTailPosition[0]--
				if currentHeadPosition[1] > currentTailPosition[1] {
					currentTailPosition[1]++
				} else {
					currentTailPosition[1]--
				}
			} else if currentHeadPosition[1]-currentTailPosition[1] == 2 {
				currentTailPosition[1]++
				if currentHeadPosition[0] > currentTailPosition[0] {
					currentTailPosition[0]++
				} else {
					currentTailPosition[0]--
				}
			} else if currentHeadPosition[1]-currentTailPosition[1] == -2 {
				currentTailPosition[1]--
				if currentHeadPosition[0] > currentTailPosition[0] {
					currentTailPosition[0]++
				} else {
					currentTailPosition[0]--
				}
			}
		}
		tailPositions = append(tailPositions, currentTailPosition)
	}
	tailPositions = append(tailPositions, currentTailPosition)
	fmt.Println("Current Head Position: ", currentHeadPosition)
	fmt.Println("Current Tail Position: ", currentTailPosition)
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		//fmt.Println("File Line", stringVal)
		stringSlice := strings.Fields(stringVal)
		direction := stringSlice[0]
		steps, _ := strconv.Atoi(stringSlice[1])
		move(direction, steps)
	}

	fmt.Println("Part 1, count of Tail Positions: ", len(getUniquePositions(tailPositions)))
}
