package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("day11/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	soroundings := []int{-11, -10, -9, -1, 1, 9, 10, 11}

	var octopuses [100]int

	c := 0
	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), "") {
			n, _ := strconv.Atoi(v)
			octopuses[c] = n
			c++
		}
	}

	totalFlashes := 0
	for i := 1; i <= 100; i++ {
		for k := range octopuses {
			octopuses[k]++
		}
		flashes := 1
		for flashes > 0 {
			flashes = 0
			for k := range octopuses {
				if octopuses[k] > 9 {
					flashes++
					totalFlashes++
					for _, v := range soroundings {
						pos := k + v
						if pos < 0 || pos >= 100 {
							continue
						}
						if octopuses[pos] == 0 {
							continue
						}
						if k < 10 {
							if v == -11 || v == -10 || v == -9 {
								continue
							}
						}
						if k > 90 {
							if v == 9 || v == 10 || v == 11 {
								continue
							}
						}
						if k%10 == 0 {
							if v == -11 || v == -1 || v == 9 {
								continue
							}
						}
						if k%10 == 9 {
							if v == -9 || v == 1 || v == 11 {
								continue
							}
						}
						octopuses[pos]++
					}
					octopuses[k] = 0
				}
			}
		}
	}

	fmt.Printf("part 1 => %d\n", totalFlashes)
}
