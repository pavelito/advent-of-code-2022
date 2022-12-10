package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSignalStrengths(values []int) []int {

	strengths := []int{values[20] * 20}
	for cycle := 21; cycle <= len(values); cycle++ {
		if (cycle-20)%40 == 0 {
			strengths = append(strengths, values[cycle]*cycle)
		}
	}

	return strengths
}

func getTotalSignalStrength(strengths []int) int {
	total := 0
	for _, v := range strengths {
		total += v
	}
	return total
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	registerValueDuring := []int{0}
	registerValueAfter := []int{0}

	//At the beginning register has a value of 1 and cycle count 0
	register := 1
	cycleCount := 1

	for fileScanner.Scan() {
		command := fileScanner.Text()
		//fmt.Println("File Line", command)

		if command == "noop" {

			//registerValueDuring[cycleCount] = register
			registerValueDuring = append(registerValueDuring, register)
			//registerValueAfter[cycleCount] = register
			registerValueAfter = append(registerValueAfter, register)
			cycleCount++

		} else {

			//registerValueDuring[cycleCount] = register
			registerValueDuring = append(registerValueDuring, register)
			//registerValueAfter[cycleCount] = register
			registerValueAfter = append(registerValueAfter, register)
			cycleCount++

			//registerValueDuring[cycleCount] = register
			registerValueDuring = append(registerValueDuring, register)

			registerIncrement, _ := strconv.Atoi(strings.Fields(command)[1])
			register = registerValueDuring[cycleCount] + registerIncrement

			//registerValueAfter[cycleCount] = register
			registerValueAfter = append(registerValueAfter, register)

			cycleCount++

		}

		// fmt.Println("Cycle Count: ", cycleCount)
		// fmt.Println("Register Value During: ", registerValueDuring)
		// fmt.Println("Register Value After: ", registerValueAfter)
	}
	signalStrengths := getSignalStrengths(registerValueDuring)
	fmt.Println("Strengths: ", signalStrengths)
	fmt.Println("Part 1, Sum of Strengths: ", getTotalSignalStrength(signalStrengths))

}
