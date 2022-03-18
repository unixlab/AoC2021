package day08

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DigitCounter [10]int

func (dc DigitCounter) Sum() int {
	sum := 0
	for _, v := range dc {
		sum += v
	}
	return sum
}

func Run() {
	file, err := os.Open("day08/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var digitCounter DigitCounter

	for scanner.Scan() {
		output := strings.Split(scanner.Text(), "|")[1]

		for _, digit := range strings.Split(output, " ") {
			switch len(digit) {
			case 2:
				digitCounter[1]++
			case 3:
				digitCounter[7]++
			case 4:
				digitCounter[4]++
			case 7:
				digitCounter[8]++
			}
		}
	}
	fmt.Printf("part 1 = %d\n", digitCounter.Sum())
}
