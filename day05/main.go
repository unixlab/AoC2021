package day05

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Run() {
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

		if x1 != x2 && y1 != y2 {
			continue
		}

		var xStart, xEnd, yStart, yEnd int
		if x1 > x2 {
			xStart = x2
			xEnd = x1
		} else {
			xStart = x1
			xEnd = x2
		}
		if y1 > y2 {
			yStart = y2
			yEnd = y1
		} else {
			yStart = y1
			yEnd = y2
		}

		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				grid[x][y]++
			}
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
	fmt.Printf("part 1 => %d\n", counter)
}
