package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoSheet struct {
	Row []BingoRow
	Sum int
}

type BingoRow struct {
	Row   []int
	Drawn int
}

func (bs BingoSheet) checkBingo() bool {
	for _, v := range bs.Row {
		if v.Drawn == 5 {
			return true
		}
	}
	return false
}

func Run() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var draws []int
	var sheet BingoSheet
	var sheets []BingoSheet

	counter := -1
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if counter == -1 {
			drawsString := strings.Split(scanner.Text(), ",")
			for _, v := range drawsString {
				number, _ := strconv.Atoi(v)
				draws = append(draws, number)
			}
			counter++
			continue
		}
		var row BingoRow
		row.Drawn = 0
		for i := 0; i < 14; i += 3 {
			number, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()[i : i+2]))
			row.Row = append(row.Row, number)
			sheet.Sum += number
		}
		sheet.Row = append(sheet.Row, row)
		if counter == 4 {
			sheets = append(sheets, sheet)
			sheet = BingoSheet{}
			counter = 0
			continue
		}
		counter++
	}

Draws:
	for _, draw := range draws {
		for i := 0; i < len(sheets); i++ {
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					if sheets[i].Row[x].Row[y] == draw {
						sheets[i].Sum -= draw
						sheets[i].Row[x].Drawn += 1
					}
				}
			}
			if sheets[i].checkBingo() {
				fmt.Printf("part 1 => %d\n", sheets[i].Sum*draw)
				break Draws
			}
		}
	}
}
