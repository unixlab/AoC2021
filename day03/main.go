package day03

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringBitsToInt(input string) int {
	returnInt := 0
	multiplier := 1
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 49 {
			returnInt += multiplier
		}
		multiplier = multiplier * 2
	}
	return returnInt
}

func Run() {
	var numbers [][]string
	file, err := os.Open("day03/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers = append(numbers, strings.Split(scanner.Text(), ""))
	}

	var gamma, epsilon strings.Builder

	for i := 0; i < len(numbers[0]); i++ {
		count0 := 0
		for _, v := range numbers {
			if v[i] == "0" {
				count0++
			}
		}
		if count0 >= (len(numbers) / 2) {
			gamma.WriteString("0")
			epsilon.WriteString("1")
		} else {
			gamma.WriteString("1")
			epsilon.WriteString("0")
		}
	}

	gammaIntValue := stringBitsToInt(gamma.String())
	epsilonIntValue := stringBitsToInt(epsilon.String())

	fmt.Printf("part 1 => %d\n", gammaIntValue*epsilonIntValue)
}
