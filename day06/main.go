package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("day06/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var fish [10]int

	for scanner.Scan() {
		for _, n := range strings.Split(scanner.Text(), ",") {
			intFish, _ := strconv.Atoi(n)
			fish[intFish]++
		}
	}

	for i := 0; i < 80; i++ {
		fish[7] += fish[0]
		fish[9] = fish[0]
		for j := 0; j <= 8; j++ {
			fish[j] = fish[j+1]
		}
	}

	fish[9] = 0

	numberOfFish := 0
	for _, v := range fish {
		numberOfFish += v
	}
	fmt.Printf("part 1 => %d\n", numberOfFish)
}
