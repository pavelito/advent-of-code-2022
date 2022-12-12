package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/gammazero/deque"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	min := math.MaxFloat64
	grid, start, end, a_list := getGridFromInput(input)

	fmt.Println("Part 1: ", bfs(grid, start, end))

	for _, a := range a_list {
		steps := float64(bfs(grid, a, end))
		min = math.Min(min, steps)
	}

	fmt.Println("Part 2: ", int(min))
}

func getGridFromInput(input []byte) ([][]int, [2]int, [2]int, [][2]int) {
	var grid [][]int
	var a_list [][2]int
	start, end := [2]int{0, 0}, [2]int{0, 0}
	for i, line := range strings.Split(string(input), "\n") {
		var row []int
		for j, v := range line {
			row = append(row, int(v))
			if v == 'S' {
				start = [2]int{i, j}
				a_list = append(a_list, [2]int{i, j})
				row[j] = int('a')
			} else if v == 'E' {
				end = [2]int{i, j}
				row[j] = int('z')
			} else if v == 'a' {
				a_list = append(a_list, [2]int{i, j})
			}
		}
		grid = append(grid, row)
	}

	return grid, start, end, a_list
}

type Node struct {
	pos  [2]int
	dist int
}

func inBound(i int, j int, grid [][]int) bool {
	if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) {
		return true
	}
	return false
}

func bfs(grid [][]int, start [2]int, end [2]int) int {
	var q deque.Deque[Node]
	q.PushBack(Node{start, 0})
	seen := make(map[[2]int]bool)
	for q.Len() > 0 {
		v := q.PopFront()
		pos, dist := v.pos, v.dist
		if pos == end {
			return dist
		}
		if seen[pos] {
			continue
		}
		seen[pos] = true
		x, y := pos[0], pos[1]
		coords := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, coord := range coords {
			dx, dy := coord[0], coord[1]
			if inBound(x+dx, y+dy, grid) && grid[x+dx][y+dy]-grid[x][y] <= 1 {
				q.PushBack(Node{[2]int{x + dx, y + dy}, dist + 1})
			}
		}

	}
	return math.MaxInt
}
