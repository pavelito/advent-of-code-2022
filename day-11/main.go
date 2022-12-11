package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var leastCommonMultiple int

type Monkey struct {
	items           []int
	inspectionCount int
	operationType   string
	operationFactor int
	testFactor      int
	testTrueMonkey  int
	testFalseMonkey int
}

func executeRound(monkeys []*Monkey, isPart2 bool) {
	for _, v := range monkeys {
		//fmt.Println("Executing Monkey ", i)
		for _, item := range v.items {
			worryLevel := item //start value
			//operation
			switch v.operationType {
			case "add":
				worryLevel = worryLevel + v.operationFactor
				break
			case "multiply":
				worryLevel = worryLevel * v.operationFactor
				break
			case "power":
				worryLevel = worryLevel * worryLevel
				break
			}
			//
			//increment the inspection count for that monkey
			v.inspectionCount++

			//Part 2
			if isPart2 {
				worryLevel = worryLevel % leastCommonMultiple
			} else {
				//gets bored now Part 1
				worryLevel = int(math.Round(float64(worryLevel / 3)))
			}

			//decide where to throw
			var targetMonkey *Monkey
			if worryLevel%v.testFactor == 0 {
				//throw to test True Monkey
				//fmt.Println("Throwing value ", worryLevel, "  to Monkey: ", v.testTrueMonkey)
				targetMonkey = monkeys[v.testTrueMonkey]

			} else {
				//throw to test False Monkey
				//fmt.Println("Throwing value ", worryLevel, "  to Monkey: ", v.testFalseMonkey)
				targetMonkey = monkeys[v.testFalseMonkey]
			}
			targetMonkey.items = append(targetMonkey.items, worryLevel)
			v.items = v.items[1:]
		}
	}
}

func getMonkeyFromLines(lines []string) *Monkey {
	// lineSlice := strings.Fields(line)
	//fmt.Println(lines)
	monkey := Monkey{}
	for _, line := range lines {
		lineSlice := strings.Fields(line)
		if len(lineSlice) == 0 {
			continue
		}
		if lineSlice[0] == "Monkey" {
			//do nothing
		} else if lineSlice[0] == "Starting" {
			//generate items list

			for _, v := range lineSlice[2:] {
				item := strings.Replace(v, ",", "", -1)
				itemInt, _ := strconv.Atoi((item))
				monkey.items = append(monkey.items, itemInt)
			}
		} else if lineSlice[0] == "Operation:" {
			operationFactor, _ := strconv.Atoi(lineSlice[5])
			//generate operation
			switch lineSlice[4] {
			case "+":
				monkey.operationType = "add"
				monkey.operationFactor = operationFactor
				break
			case "*":
				if lineSlice[5] == "old" {
					monkey.operationType = "power"
					monkey.operationFactor = 0
				} else {
					monkey.operationType = "multiply"
					monkey.operationFactor = operationFactor
				}
				break
			}
		} else if lineSlice[0] == "Test:" {
			testValue, _ := strconv.Atoi(lineSlice[3])
			monkey.testFactor = testValue
		} else if lineSlice[0] == "If" {
			monkeyValue, _ := strconv.Atoi(lineSlice[5])
			if lineSlice[1] == "true:" {
				monkey.testTrueMonkey = monkeyValue
			} else {
				monkey.testFalseMonkey = monkeyValue
			}
		}
	}
	monkey.inspectionCount = 0
	return &monkey
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	monkeys := []*Monkey{}
	monkeyLines := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		//fmt.Println("File Line", stringVal)

		monkeyLines = append(monkeyLines, line)
		// fmt.Println("Line: ", line)
		// fmt.Println("Lines: ", monkeyLines)
		if line == "" {
			monkeys = append(monkeys, getMonkeyFromLines(monkeyLines))
			monkeyLines = []string{}
		}
	}

	leastCommonMultiple = 1
	for _, m := range monkeys {
		leastCommonMultiple = leastCommonMultiple * m.testFactor
	}

	//Begin Logic Part 1
	for round := 1; round <= 20; round++ {
		//fmt.Println("Executing Round: ", round)
		executeRound(monkeys, false)
		if round == 1 || round == 20 || round == 1000 || round == 10000 {
			fmt.Println("After Round: ", round)
			for _, v := range monkeys {
				fmt.Println("Monkey Activity: ", v.inspectionCount)
			}
			fmt.Println()
		}
	}

	inspectionCounts := []int{}
	for _, v := range monkeys {
		inspectionCounts = append(inspectionCounts, v.inspectionCount)
	}
	sort.Ints(inspectionCounts)
	fmt.Println("Part 1: ", inspectionCounts[len(inspectionCounts)-1]*inspectionCounts[len(inspectionCounts)-2])

	//Begin Logic Part 2
	for round := 1; round <= 10000; round++ {
		//fmt.Println("Executing Round: ", round)
		executeRound(monkeys, true)
		if round == 1 || round == 20 || round == 1000 || round == 10000 {
			fmt.Println("After Round: ", round)
			for _, v := range monkeys {
				fmt.Println("Monkey Activity: ", v.inspectionCount)
			}
			fmt.Println()
		}
	}

	inspectionCounts = []int{}
	for _, v := range monkeys {
		inspectionCounts = append(inspectionCounts, v.inspectionCount)
	}
	sort.Ints(inspectionCounts)
	fmt.Println("Part 2: ", inspectionCounts[len(inspectionCounts)-1]*inspectionCounts[len(inspectionCounts)-2])

}
