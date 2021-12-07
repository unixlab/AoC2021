package day07

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func intAbs(number int) int {
	return int(math.Abs(float64(number)))
}

func Run() {
	file, err := os.Open("day07/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var positions []int

	for scanner.Scan() {
		for _, n := range strings.Split(scanner.Text(), ",") {
			pos, _ := strconv.Atoi(n)
			positions = append(positions, pos)
		}
	}

	minPos := -1
	maxPos := -1
	for _, v := range positions {
		if v < minPos || minPos == -1 {
			minPos = v
		}
		if v > maxPos {
			maxPos = v
		}
	}

	minFuelP1 := -1
	minFuelP2 := -1
	for i := minPos; i < maxPos; i++ {
		posFuelP1 := 0
		posFuelP2 := 0
		for _, v := range positions {
			if minFuelP1 > 0 && posFuelP1 > minFuelP1 && minFuelP2 > 0 && posFuelP2 > minFuelP2 {
				break
			}
			distance := intAbs(v - i)
			posFuelP1 += distance
			posFuelP2 += distance * (distance + 1) / 2
		}
		if posFuelP1 < minFuelP1 || minFuelP1 == -1 {
			minFuelP1 = posFuelP1
		}
		if posFuelP2 < minFuelP2 || minFuelP2 == -1 {
			minFuelP2 = posFuelP2
		}
	}
	fmt.Printf("part 1 => %d\n", minFuelP1)
	fmt.Printf("part 2 => %d\n", minFuelP2)
}
