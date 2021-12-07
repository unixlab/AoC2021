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

	minFuel := -1
	for i := minPos; i < maxPos; i++ {
		posFuel := 0
		for _, v := range positions {
			if minFuel > 0 && posFuel > minFuel {
				break
			}
			posFuel += intAbs(v - i)
		}
		if posFuel < minFuel || minFuel == -1 {
			minFuel = posFuel
		}
	}
	fmt.Printf("part 1 => %d\n", minFuel)
}
