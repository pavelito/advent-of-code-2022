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
var part1TailPositions [][2]int

var currentPositions map[string][2]int

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

func getHeadPosition(direction string, currentPosition [2]int) [2]int {
	switch direction {
	case "R":
		currentPosition[1] = currentPosition[1] + 1
	case "U":
		currentPosition[0] = currentPosition[0] + 1
	case "L":
		currentPosition[1] = currentPosition[1] - 1
	case "D":
		currentPosition[0] = currentPosition[0] - 1
	}
	return currentPosition
}

func getTailPosition(currentHeadPosition [2]int, currentTailPosition [2]int) [2]int {
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
	// tailPositions = append(tailPositions, currentTailPosition)
	return currentTailPosition
}

func move(direction string, steps int) {
	//fmt.Println("Moving ", direction, steps)

	for step := 1; step <= steps; step++ {
		//move head
		currentPositions["H"] = getHeadPosition(direction, currentPositions["H"])
		headPositions = append(headPositions, currentPositions["H"])

		//move intermediate knots
		currentPositions["1"] = getTailPosition(currentPositions["H"], currentPositions["1"])
		currentPositions["2"] = getTailPosition(currentPositions["1"], currentPositions["2"])
		currentPositions["3"] = getTailPosition(currentPositions["2"], currentPositions["3"])
		currentPositions["4"] = getTailPosition(currentPositions["3"], currentPositions["4"])
		currentPositions["5"] = getTailPosition(currentPositions["4"], currentPositions["5"])
		currentPositions["6"] = getTailPosition(currentPositions["5"], currentPositions["6"])
		currentPositions["7"] = getTailPosition(currentPositions["6"], currentPositions["7"])
		currentPositions["8"] = getTailPosition(currentPositions["7"], currentPositions["8"])

		//move tail now
		currentPositions["9"] = getTailPosition(currentPositions["8"], currentPositions["9"])
		tailPositions = append(tailPositions, currentPositions["9"])

		//track part 1 tail positions
		part1TailPositions = append(part1TailPositions, currentPositions["1"])

	}
	//fmt.Println("Current Positions: ", currentPositions)
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	currentPositions = map[string][2]int{}
	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		//fmt.Println("File Line", stringVal)
		stringSlice := strings.Fields(stringVal)
		direction := stringSlice[0]
		steps, _ := strconv.Atoi(stringSlice[1])
		move(direction, steps)
	}

	fmt.Println("Part 1, count of Tail Positions: ", len(getUniquePositions(part1TailPositions)))
	fmt.Println("Part 2, count of Tail Positions: ", len(getUniquePositions(tailPositions)))
}
