package day05

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func RunPart(part int) {
	file, err := os.Open("day05/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	regex, _ := regexp.Compile("^(\\d+),(\\d+) -> (\\d+),(\\d+)$")

	var grid [1000][1000]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = 0
		}
	}

	for scanner.Scan() {
		cords := regex.FindAllStringSubmatch(scanner.Text(), -1)

		x1, _ := strconv.Atoi(cords[0][1])
		y1, _ := strconv.Atoi(cords[0][2])
		x2, _ := strconv.Atoi(cords[0][3])
		y2, _ := strconv.Atoi(cords[0][4])

		if part == 1 {
			if x1 != x2 && y1 != y2 {
				continue
			}
		}

		var xOperator, yOperator int

		if x1 > x2 {
			xOperator = -1
		} else {
			xOperator = 1
		}
		if y1 > y2 {
			yOperator = -1
		} else {
			yOperator = 1
		}

		x := x1
		for {
			if xOperator == 1 && x > x2 {
				break
			}
			if xOperator == -1 && x < x2 {
				break
			}
			y := y1
			for {
				if yOperator == 1 && y > y2 {
					break
				}
				if yOperator == -1 && y < y2 {
					break
				}
				grid[x][y]++
				y += yOperator
				if x1 != x2 && y1 != y2 {
					x += xOperator
				}
			}
			x += xOperator
		}
	}

	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				counter++
			}
		}
	}
	fmt.Printf("part %d => %d\n", part, counter)
}
