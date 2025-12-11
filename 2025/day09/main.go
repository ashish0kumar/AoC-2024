package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pt struct{ x, y int }
type seg struct{ a, b pt }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func area(a, b pt) int {
	return (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
}

func readFile(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func parse(lines []string) []pt {
	var ps []pt
	for _, s := range lines {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		parts := strings.Split(s, ",")
		if len(parts) < 2 {
			continue
		}

		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		ps = append(ps, pt{x, y})
	}
	return ps
}

func buildSegs(p []pt) []seg {
	n := len(p)
	var ss []seg

	for i := 0; i < n-1; i++ {
		ss = append(ss, seg{p[i], p[i+1]})
	}
	ss = append(ss, seg{p[n-1], p[0]})
	return ss
}

func (s *seg) intersects(r1, r2 pt) bool {
	rminx := min(r1.x, r2.x) + 1
	rmaxx := max(r1.x, r2.x) - 1
	rminy := min(r1.y, r2.y) + 1
	rmaxy := max(r1.y, r2.y) - 1

	sminx := min(s.a.x, s.b.x)
	smaxx := max(s.a.x, s.b.x)
	sminy := min(s.a.y, s.b.y)
	smaxy := max(s.a.y, s.b.y)

	if smaxx < rminx || sminx > rmaxx {
		return false
	}
	if smaxy < rminy || sminy > rmaxy {
		return false
	}
	return true
}

func solve(part2 bool, lines []string) int {
	p := parse(lines)
	n := len(p)
	if n == 0 {
		return 0
	}

	if !part2 {
		best := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				a := area(p[i], p[j])
				if a > best {
					best = a
				}
			}
		}
		return best
	}

	ss := buildSegs(p)
	best := 0
	for i := 0; i < n-1; i++ {
	main:
		for j := i + 1; j < n; j++ {
			a := area(p[i], p[j])
			if a <= best {
				continue
			}
			for k := range ss {
				if ss[k].intersects(p[i], p[j]) {
					continue main
				}
			}
			best = a
		}
	}
	return best
}

func main() {
	lines := readFile("input.txt")
	fmt.Println("Part1:", solve(false, lines))
	fmt.Println("Part2:", solve(true, lines))
}