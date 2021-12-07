package dayXX

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	file, err := os.Open("dayXX/example-input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
