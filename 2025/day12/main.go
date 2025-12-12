package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type grid [][]rune
type pres struct {
	shape grid
	area  int
}
type region struct {
	x, y   int
	cnts []int
}

func readFile(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var out []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		out = append(out, sc.Text())
	}
	return out
}

func parse(lines []string) ([6]pres, []region) {
	var ps [6]pres
	var rs []region
	cur := 0

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		if strings.Contains(l, ":") && !strings.Contains(l, "x") {
			parts := strings.SplitN(l, ":", 2)
			n, _ := strconv.Atoi(parts[0])
			cur = n

		} else if strings.Contains(l, "#") {
			ps[cur].shape = append(ps[cur].shape, []rune(l))

		} else if strings.Contains(l, "x") {
			parts := strings.SplitN(l, ":", 2)
			dims := strings.Split(parts[0], "x")

			w, _ := strconv.Atoi(dims[0])
			h, _ := strconv.Atoi(dims[1])

			fields := strings.Fields(parts[1])
			cnt := make([]int, len(fields))

			for i, f := range fields {
				v, _ := strconv.Atoi(f)
				cnt[i] = v
			}
			rs = append(rs, region{x: w, y: h, cnts: cnt})
		}
	}
	return ps, rs
}

func calcAreas(ps *[6]pres) {
	for i := range ps {
		a := 0
		for _, r := range ps[i].shape {
			for _, c := range r {
				if c == '#' {
					a++
				}
			}
		}
		ps[i].area = a
	}
}

func solve(lines []string) int {
	ps, rs := parse(lines)
	calcAreas(&ps)
	res := 0
	
	for _, r := range rs {
		total := r.x * r.y
		sum := 0

		for i, c := range r.cnts {
			if i < len(ps) {
				sum += ps[i].area * c
			}
		}
		if sum < total {
			res++
		}
	}
	return res
}

func main() {
	lines := readFile("input.txt")
	fmt.Println("Part 1:", solve(lines))
}