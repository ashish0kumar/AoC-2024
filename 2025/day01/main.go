package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func solve(part2 bool, inputs []string) int {
	var dir int
	pos := 50
	cnt := 0

	for _, line := range inputs {
		prev := pos
		direction := line[0]

		switch direction {
		case 76:
			dir = -1;
		case 82:
			dir = 1;
		}

		dist, _ := strconv.Atoi(line[1:])
		if part2 && dist >= 100 {
			cnt += dist / 100
		}

		pos += (dist % 100) * dir
		if pos > 99 {
			pos -= 100
			if part2 && pos != 0 && prev != 0 {
				cnt++
			}
		}

		if pos < 0 {
			pos += 100
			if part2 && pos != 0 && prev != 0 {
				cnt++
			}
		}
		if pos == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {

	content := readFile("input.txt")
	sumP1 := solve(false, content)
	sumP2 := solve(true, content)

	fmt.Println("Part1:", sumP1)
	fmt.Println("Part2:", sumP2)
}
