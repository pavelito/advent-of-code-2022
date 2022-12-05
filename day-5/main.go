package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func getStacks() []Stack {

	//Trial
	// 	   [D]
	// [N] [C]
	// [Z] [M] [P]
	// 1   2   3
	// return []Stack{
	// 	{"Z", "M"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }

	//Real
	// [N]     [Q]         [N]
	// [R]     [F] [Q]     [G] [M]
	// [J]     [Z] [T]     [R] [H] [J]
	// [T] [H] [G] [R]     [B] [N] [T]
	// [Z] [J] [J] [G] [F] [Z] [S] [M]
	// [B] [N] [N] [N] [Q] [W] [L] [Q] [S]
	// [D] [S] [R] [V] [T] [C] [C] [N] [G]
	// [F] [R] [C] [F] [L] [Q] [F] [D] [P]
	// 1   2   3   4   5   6   7   8   9
	return []Stack{
		{"F", "D", "B", "Z", "T", "J", "R", "N"},
		{"R", "S", "N", "J", "H"},
		{"C", "R", "N", "J", "G", "Z", "F", "Q"},
		{"F", "V", "N", "G", "R", "T", "Q"},
		{"L", "T", "Q", "F"},
		{"Q", "C", "W", "Z", "B", "R", "G", "N"},
		{"F", "C", "L", "S", "N", "H", "M"},
		{"D", "N", "Q", "M", "T", "J"},
		{"P", "G", "S"},
	}
}

type move struct {
	count int
	from  int
	to    int
}

func decodeMove(moveText string) move {
	moveArray := strings.Fields(moveText)
	moveFrom, _ := strconv.Atoi(moveArray[3])
	moveTo, _ := strconv.Atoi(moveArray[5])
	moveCount, _ := strconv.Atoi(moveArray[1])
	return move{
		count: moveCount,
		from:  moveFrom - 1,
		to:    moveTo - 1,
	}
}

func operateOnStack(stacks []Stack, instruction move) {
	for i := 1; i <= instruction.count; i++ {
		poppedCrate, popSuccess := stacks[instruction.from].Pop()
		if popSuccess {
			stacks[instruction.to].Push(poppedCrate)
		}
	}
}

func main() {

	stacksOfCrates := getStacks()

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		stringVal := fileScanner.Text()
		moveInstructions := decodeMove(stringVal)
		//fmt.Println(moveInstructions)
		//fmt.Println("File Line", stringVal)
		operateOnStack(stacksOfCrates, moveInstructions)
		fmt.Println(stacksOfCrates)
	}
	fmt.Println("Final Result")
	for _, s := range stacksOfCrates {
		fmt.Println(s)
	}

	// fmt.Println(stackA)

	readFile.Close()
}
