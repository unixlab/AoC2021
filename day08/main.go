package day08

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
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

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	sum := 0

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		signal := strings.Split(scanner.Text(), "|")[0]
		output := strings.Split(scanner.Text(), "|")[1]

		var digits [10]string

		// get 1, 4, 7 and 8
		for _, digit := range strings.Split(signal, " ") {
			switch len(digit) {
			case 2:
				digits[1] = digit
			case 4:
				digits[4] = digit
			case 3:
				digits[7] = digit
			case 7:
				digits[8] = digit
			}
		}

		// get 3
		for _, digit := range strings.Split(signal, " ") {
			if len(digit) != 5 {
				continue
			}
			if strings.Index(digit, digits[1][0:1]) < 0 {
				continue
			}
			if strings.Index(digit, digits[1][1:2]) < 0 {
				continue
			}
			digits[3] = digit
		}

		// get 9
	Digit9:
		for _, digit := range strings.Split(signal, " ") {
			if len(digit) != 6 {
				continue
			}
			for _, v := range strings.Split(digits[3], "") {
				if strings.Index(digit, v) < 0 {
					continue Digit9
				}
			}
			for _, v := range strings.Split(digits[4], "") {
				if strings.Index(digit, v) < 0 {
					continue Digit9
				}
			}
			digits[9] = digit
		}

		// get 0
	Digit0:
		for _, digit := range strings.Split(signal, " ") {
			for _, v := range digits {
				if digit == v {
					continue Digit0
				}
			}
			if len(digit) != 6 {
				continue
			}
			if strings.Index(digit, digits[1][0:1]) < 0 {
				continue
			}
			if strings.Index(digit, digits[1][1:2]) < 0 {
				continue
			}
			digits[0] = digit
		}

		// get 6
	Digit6:
		for _, digit := range strings.Split(signal, " ") {
			for _, v := range digits {
				if digit == v {
					continue Digit6
				}
			}
			if len(digit) != 6 {
				continue
			}
			digits[6] = digit
		}

		twoIndicator := ""
		if strings.Index(digits[6], digits[1][0:1]) < 0 {
			twoIndicator = digits[1][1:2]
		}
		if strings.Index(digits[6], digits[1][1:2]) < 0 {
			twoIndicator = digits[1][0:1]
		}

		// get 5
	Digit5:
		for _, digit := range strings.Split(signal, " ") {
			for _, v := range digits {
				if digit == v {
					continue Digit5
				}
			}
			if len(digit) != 5 {
				continue
			}
			if strings.Index(digit, twoIndicator) < 0 {
				continue
			}
			digits[5] = digit
		}

		// get 2
	Digit2:
		for _, digit := range strings.Split(signal, " ") {
			for _, v := range digits {
				if digit == v {
					continue Digit2
				}
			}
			if len(digit) != 5 {
				continue
			}
			digits[2] = digit
		}

		number := 0
		for k, o := range strings.Split(output, " ") {
			for n, d := range digits {
				if SortString(o) == SortString(d) {
					for i := 4; i > k; i-- {
						n = n * 10
					}
					number += n
				}
			}
		}

		sum += number
	}

	fmt.Printf("part 2 = %d\n", sum)
}
