package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	inputFile, err := os.Open("day01/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	increasedP1 := 0
	increasedP2 := 0
	lastMeasure := -1

	n1, n2, n3 := 0, 0, 0

	for scanner.Scan() {
		number, parseError := strconv.Atoi(scanner.Text())
		if parseError != nil {
			panic(parseError)
		}

		if lastMeasure > 0 && number > lastMeasure {
			increasedP1++
		}

		lastMeasure = number

		if n1 > 0 && n2 > 2 && n3 > 0 {
			if number+n1+n2 > n1+n2+n3 {
				increasedP2++
			}
		}

		n3 = n2
		n2 = n1
		n1 = number
	}

	fmt.Printf("part 1 => %d\n", increasedP1)
	fmt.Printf("part 2 => %d\n", increasedP2)
}
