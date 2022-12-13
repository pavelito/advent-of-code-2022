package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	messages := [][]any{}
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		var packet []any
		json.Unmarshal([]byte(line), &packet)
		messages = append(messages, packet)

	}

	// Part 1
	sumOfIndexes := 0
	for i := 0; i < len(messages); i += 2 {
		if isLess(messages[i], messages[i+1]) {
			sumOfIndexes += (i / 2) + 1
		}
	}
	fmt.Println("Part 1: ", sumOfIndexes)

	// Part 2
	// Add divider packets
	var divider1 []any
	var divider2 []any
	json.Unmarshal([]byte("[[2]]"), &divider1)
	json.Unmarshal([]byte("[[6]]"), &divider2)
	messages = append(messages, divider1, divider2)

	// Sort packets
	sort.Slice(
		messages,
		func(i, j int) bool {
			return isLess(messages[i], messages[j])
		},
	)

	// Compute decoder key
	decoderKey := 1
	for i, message := range messages {
		if compare(message, divider1) == 0 || compare(message, divider2) == 0 {
			decoderKey *= i + 1
		}
	}
	fmt.Println("Part 2: ", decoderKey)
}

func isLess(message1 []any, message2 []any) bool {
	return compare(message1, message2) < 1
}

func compare(message1 []any, message2 []any) int {
	for i := 0; i < len(message1) && i < len(message2); i++ {
		l := message1[i]
		r := message2[i]

		lIsNum := reflect.TypeOf(l).Name() == "float64"
		rIsNum := reflect.TypeOf(r).Name() == "float64"

		if lIsNum && rIsNum {
			if l.(float64) < r.(float64) {
				return -1
			}
			if l.(float64) > r.(float64) {
				return 1
			}
		} else {
			var lChildren []any
			var rChildren []any

			if lIsNum {
				lChildren = []any{l}
			} else {
				lChildren = l.([]any)
			}

			if rIsNum {
				rChildren = []any{r}
			} else {
				rChildren = r.([]any)
			}

			res := compare(lChildren, rChildren)

			if res != 0 {
				return res
			}
		}
	}

	if len(message1) < len(message2) {
		return -1
	}

	if len(message1) > len(message2) {
		return 1
	}

	return 0
}
