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

type pt struct {
	r, c int
}

func solve(part2 bool, lines []string) int64 {
	h := len(lines)
	if h == 0 {
		return 0
	}

	g := make([][]byte, h)
	for i := 0; i < h; i++ {
		g[i] = []byte(lines[i])
	}
	w := len(g[0])

	sr, sc := -1, -1
	for i := 0; i < h; i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == 'S' {
				sr, sc = i, j
				break
			}
		}
		if sr != -1 {
			break
		}
	}
	if sr == -1 || sr+1 >= h {
		if part2 {
			return 1
		}
		return 0
	}

	if !part2 {
		cur := make(map[pt]bool)
		cur[pt{sr + 1, sc}] = true
		seen := make(map[pt]bool)
		var splits int64

		for len(cur) > 0 {
			nxt := make(map[pt]bool)

			for p := range cur {
				r, c := p.r, p.c
				if r < 0 || r >= h || c < 0 || c >= w {
					continue
				}
				ch := g[r][c]

				if ch == '^' {
					if !seen[p] {
						seen[p] = true
						splits++
					}
					if c-1 >= 0 {
						nxt[pt{r, c - 1}] = true
					}
					if c+1 < w {
						nxt[pt{r, c + 1}] = true
					}
				} else {
					nr := r + 1
					if nr < h {
						nxt[pt{nr, c}] = true
					}
				}
			}

			cur = nxt
		}

		return splits
	}

	cur := make(map[pt]int64)
	cur[pt{sr + 1, sc}] = 1
	var timelines int64

	for len(cur) > 0 {
		nxt := make(map[pt]int64)

		for p, cnt := range cur {
			r, c := p.r, p.c
			if r < 0 || r >= h || c < 0 || c >= w {
				timelines += cnt
				continue
			}
			ch := g[r][c]

			if ch == '^' {
				if c-1 >= 0 {
					nxt[pt{r, c - 1}] += cnt
				} else {
					timelines += cnt
				}
				if c+1 < w {
					nxt[pt{r, c + 1}] += cnt
				} else {
					timelines += cnt
				}
			} else {
				nr := r + 1
				if nr < h {
					nxt[pt{nr, c}] += cnt
				} else {
					timelines += cnt
				}
			}
		}

		cur = nxt
	}

	return timelines
}

func main() {
	lines := readFile("input.txt")
	fmt.Println("Part1:", solve(false, lines))
	fmt.Println("Part2:", solve(true, lines))
}