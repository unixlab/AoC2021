package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid [100][100]int

func (g Grid) checkAllNeighborsAreLower(y, x int) bool {
	for i := y - 1; i <= y+1; i++ {
		if i < 0 || i >= len(g) {
			continue
		}
		for j := x - 1; j <= x+1; j++ {
			if j < 0 || j >= len(g[i]) {
				continue
			}
			if i == y && j == x {
				continue
			}
			if i != y && j != x {
				continue
			}
			if g[i][j] <= g[y][x] {
				return false
			}
		}
	}
	return true
}

func Run() {
	file, err := os.Open("day09/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var grid Grid

	x := 0
	y := 0
	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), "") {
			n, _ := strconv.Atoi(v)
			grid[y][x] = n
			x++
		}
		y++
		x = 0
	}

	sum := 0
	for y = 0; y < len(grid); y++ {
		for x = 0; x < len(grid[y]); x++ {
			if grid.checkAllNeighborsAreLower(y, x) {
				sum += grid[y][x] + 1
			}
		}
	}
	fmt.Printf("part 1 => %d\n", sum)
}
