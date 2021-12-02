package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	horizontal := 0
	depth := 0

	for scanner.Scan() {
		increment, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])
		switch strings.Split(scanner.Text(), " ")[0] {
		case "forward":
			horizontal += increment
		case "up":
			depth -= increment
		case "down":
			depth += increment
		default:
			panic("invalid input")
		}
	}

	fmt.Printf("part 1 => %d\n", horizontal*depth)
}