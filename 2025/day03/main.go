package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(fname string) []string {
	var lines []string
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func best2(s string) int64 {
	best := -1
	maxR := -1

	for i := len(s) - 1; i >= 0; i-- {
		d := int(s[i] - '0')

		if maxR != -1 {
			v := d*10 + maxR
			if v > best {
				best = v
			}
		}
		if d > maxR {
			maxR = d
		}
	}

	if best < 0 {
		return 0
	}
	return int64(best)
}

func best12(s string) int64 {
	k := 12
	n := len(s)
	if n < k {
		return 0
	}

	rem := n - k
	st := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		d := s[i]
		for rem > 0 && len(st) > 0 && st[len(st)-1] < d {
			st = st[:len(st)-1]
			rem--
		}
		st = append(st, d)
	}
	if len(st) > k {
		st = st[:k]
	}

	var v int64
	for _, c := range st {
		v = v*10 + int64(c - '0')
	}
	return v
}

func solve(part2 bool, banks []string) int64 {
	var total int64

	for _, s := range banks {
		if len(s) < 2 {
			continue
		}

		if !part2 {
			total += best2(s)
		} else {
			total += best12(s)
		}
	}

	return total
}

func main() {
	lines := readFile("input.txt")
	fmt.Println("Part1:", solve(false, lines))
	fmt.Println("Part2:", solve(true, lines))
}