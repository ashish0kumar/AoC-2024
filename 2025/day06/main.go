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

func solve(part2 bool, lines []string) int64 {
	h := len(lines)
	if h == 0 {
		return 0
	}

	w := 0
	for _, s := range lines {
		if len(s) > w {
			w = len(s)
		}
	}

	g := make([][]byte, h)
	for i := 0; i < h; i++ {
		r := []byte(lines[i])
		if len(r) < w {
			tmp := make([]byte, w)
			copy(tmp, r)
			for j := len(r); j < w; j++ {
				tmp[j] = ' '
			}
			r = tmp
		}
		g[i] = r
	}

	opRow := h - 1

	isSep := func(c int) bool {
		for r := 0; r < h; r++ {
			if g[r][c] != ' ' {
				return false
			}
		}
		return true
	}

	var total int64
	c := 0

	for c < w {
		for c < w && isSep(c) {
			c++
		}
		if c >= w {
			break
		}
		l := c
		for c < w && !isSep(c) {
			c++
		}
		r := c - 1

		op := byte(0)
		for j := l; j <= r; j++ {
			if g[opRow][j] == '+' || g[opRow][j] == '*' {
				op = g[opRow][j]
				break
			}
		}
		if op == 0 {
			continue
		}

		var nums []int64

		if !part2 {
			for i := 0; i < opRow; i++ {
				j := l
				for j <= r && (g[i][j] < '0' || g[i][j] > '9') {
					j++
				}
				if j > r {
					continue
				}
				var v int64
				for j <= r && g[i][j] >= '0' && g[i][j] <= '9' {
					v = v*10 + int64(g[i][j]-'0')
					j++
				}
				nums = append(nums, v)
			}
		} else {
			for j := l; j <= r; j++ {
				var v int64
				has := false
				for i := 0; i < opRow; i++ {
					ch := g[i][j]
					if ch >= '0' && ch <= '9' {
						v = v*10 + int64(ch-'0')
						has = true
					}
				}
				if has {
					nums = append(nums, v)
				}
			}
		}

		if len(nums) == 0 {
			continue
		}

		var res int64
		if op == '+' {
			for _, x := range nums {
				res += x
			}
		} else {
			res = 1
			for _, x := range nums {
				res *= x
			}
		}

		total += res
	}

	return total
}

func main() {
	lines := readFile("input.txt")
	fmt.Println("Part1:", solve(false, lines))
	fmt.Println("Part2:", solve(true, lines))
}