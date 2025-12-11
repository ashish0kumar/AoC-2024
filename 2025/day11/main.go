package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type g map[string][]string

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

func parse(lines []string) g {
	m := make(g)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		parts := strings.SplitN(l, ":", 2)
		if len(parts) < 2 {
			continue
		}
		m[strings.TrimSpace(parts[0])] = strings.Fields(strings.TrimSpace(parts[1]))
	}
	return m
}

func dfsCount(gr g, u, dst string, memo map[string]int) int {
	if v, ok := memo[u]; ok {
		return v
	}
	if u == dst {
		return 1
	}

	sum := 0
	for _, v := range gr[u] {
		sum += dfsCount(gr, v, dst, memo)
	}
	memo[u] = sum
	return sum
}

func dfsReq(gr g, u, dst string, req map[string]int, mask int, memo map[string]map[int]int) int {
	if memo[u] == nil {
		memo[u] = make(map[int]int)
	}
	if v, ok := memo[u][mask]; ok {
		return v
	}

	nm := mask
	if idx, ok := req[u]; ok {
		nm |= 1 << idx
	}
	if u == dst {
		if nm == (1<<len(req))-1 {
			return 1
		}
		return 0
	}
	
	sum := 0
	for _, v := range gr[u] {
		sum += dfsReq(gr, v, dst, req, nm, memo)
	}
	memo[u][mask] = sum
	return sum
}

func solve(part2 bool, lines []string) int {
	gr := parse(lines)
	if !part2 {
		return dfsCount(gr, "you", "out", make(map[string]int))
	}
	req := map[string]int{"dac": 0, "fft": 1}
	return dfsReq(gr, "svr", "out", req, 0, make(map[string]map[int]int))
}

func main() {
	lines := readFile("input.txt")
	fmt.Println("Part1:", solve(false, lines))
	fmt.Println("Part2:", solve(true, lines))
}