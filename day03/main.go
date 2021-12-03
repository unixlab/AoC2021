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

func arrayToString(array []string) string {
	var newString strings.Builder
	for _, v := range array {
		newString.WriteString(v)
	}
	return newString.String()
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

	oxygen := make(map[string][]string, len(numbers))
	co2 := make(map[string][]string, len(numbers))
	for _, v := range numbers {
		oxygen[arrayToString(v)] = v
		co2[arrayToString(v)] = v
	}

	i := 0
	for len(oxygen) > 1 {
		var zeros, ones []string
		for k, v := range oxygen {
			if v[i] == "0" {
				zeros = append(zeros, k)
			} else {
				ones = append(ones, k)
			}
		}
		if len(ones) >= len(zeros) {
			for _, v := range zeros {
				delete(oxygen, v)
			}
		} else {
			for _, v := range ones {
				delete(oxygen, v)
			}
		}
		i++
	}

	var oxygenIntValue int
	for k, _ := range oxygen {
		oxygenIntValue = stringBitsToInt(k)
	}

	i = 0
	for len(co2) > 1 {
		var zeros, ones []string
		for k, v := range co2 {
			if v[i] == "0" {
				zeros = append(zeros, k)
			} else {
				ones = append(ones, k)
			}
		}
		if len(ones) < len(zeros) {
			for _, v := range zeros {
				delete(co2, v)
			}
		} else {
			for _, v := range ones {
				delete(co2, v)
			}
		}
		i++
	}

	var co2IntValue int
	for k, _ := range co2 {
		co2IntValue = stringBitsToInt(k)
	}

	fmt.Printf("part 1 => %d\n", oxygenIntValue*co2IntValue)
}
