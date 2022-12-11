package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items           []int
	inspectionCount int
	operationType   string
	operationFactor int
	testFactor      int
	testTrueMonkey  int
	testFalseMonkey int
}

func executeRound(monkeys []*Monkey) {
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
			//gets bored now
			worryLevel = int(math.Round(float64(worryLevel / 3)))

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
	//Begin Logic
	for round := 1; round <= 20; round++ {
		fmt.Println("Executing Round: ", round)
		executeRound(monkeys)
	}

	for _, v := range monkeys {
		fmt.Println("Monkey Activity: ", v.inspectionCount)
	}
}
