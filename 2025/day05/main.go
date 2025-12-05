package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rng struct {
	l, r int64
}

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

func parse(lines []string) ([]rng, []int64) {
	var rs []rng
	var ids []int64
	i := 0

	for i < len(lines) && lines[i] != "" {
		parts := strings.Split(lines[i], "-")
		if len(parts) == 2 {
			l, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
			r, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
			if err1 == nil && err2 == nil {
				rs = append(rs, rng{l, r})
			}
		}
		i++
	}

	if i < len(lines) && lines[i] == "" {
		i++
	}

	for i < len(lines) {
		s := strings.TrimSpace(lines[i])
		if s != "" {
			v, err := strconv.ParseInt(s, 10, 64)
			if err == nil {
				ids = append(ids, v)
			}
		}
		i++
	}

	return rs, ids
}

func solve(part2 bool, lines []string) int64 {
	rs, ids := parse(lines)

	if !part2 {
		var cnt int64
		for _, id := range ids {
			ok := false
			for _, r := range rs {
				if id >= r.l && id <= r.r {
					ok = true
					break
				}
			}
			if ok {
				cnt++
			}
		}
		return cnt
	}

	if len(rs) == 0 {
		return 0
	}

	sort.Slice(rs, func(i, j int) bool {
		if rs[i].l == rs[j].l {
			return rs[i].r < rs[j].r
		}
		return rs[i].l < rs[j].l
	})

	var merged []rng
	cur := rs[0]
	for i := 1; i < len(rs); i++ {
		r := rs[i]
		if r.l <= cur.r+1 {
			if r.r > cur.r {
				cur.r = r.r
			}
		} else {
			merged = append(merged, cur)
			cur = r
		}
	}
	merged = append(merged, cur)

	var total int64
	for _, r := range merged {
		total += r.r - r.l + 1
	}
	return total
}

func main() {
	lines := readFile("input.txt")
	fmt.Println("Part1:", solve(false, lines))
	fmt.Println("Part2:", solve(true, lines))
}