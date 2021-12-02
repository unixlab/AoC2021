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
	depthP1 := 0

	aim := 0
	depthP2 := 0

	for scanner.Scan() {
		increment, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])
		switch strings.Split(scanner.Text(), " ")[0] {
		case "forward":
			horizontal += increment
			depthP2 += aim * increment
		case "up":
			depthP1 -= increment
			aim -= increment
		case "down":
			depthP1 += increment
			aim += increment
		default:
			panic("invalid input")
		}
	}

	fmt.Printf("part 1 => %d\n", horizontal*depthP1)
	fmt.Printf("part 2 => %d\n", horizontal*depthP2)
}
