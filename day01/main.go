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

	increased := 0
	lastMeasure := -1

	for scanner.Scan() {
		number, parseError := strconv.Atoi(scanner.Text())
		if parseError != nil {
			panic(parseError)
		}

		if lastMeasure > 0 && number > lastMeasure {
			increased++
		}

		lastMeasure = number
	}

	fmt.Printf("part 1 => %d\n", increased)
}
