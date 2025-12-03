package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
		lines = strings.Split(line, ",")
	}
	return lines
}

func remDup[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, i := range sliceList {
		if _, val := allKeys[i]; !val {
			allKeys[i] = true
			list = append(list, i)
		}
	}
	return list
}

func genInvalidIDs(L, d int) []int64 {
    repeatCnt := L / d
    if repeatCnt < 2 {
        return nil
    }

    var ids []int64
    start := int64(math.Pow10(d - 1))
    end := int64(math.Pow10(d))

    for prefix := start; prefix < end; prefix++ {
        s := strconv.FormatInt(prefix, 10)
        full := strings.Repeat(s, repeatCnt)
        num, _ := strconv.ParseInt(full, 10, 64)
        ids = append(ids, num)
    }
    return ids
}

func solve(part2 bool, inputs []string) int64 {
    var total int64

    for _, input := range inputs {
        parts := strings.Split(strings.TrimSpace(input), "-")
        if len(parts) != 2 {
            continue
        }

        start, _ := strconv.ParseInt(parts[0], 10, 64)
        end, _ := strconv.ParseInt(parts[1], 10, 64)
        startLen := len(parts[0])
        endLen := len(parts[1])

        var invalid []int64

        for L := startLen; L <= endLen; L++ {
            if !part2 {
                if L % 2 != 0 {
                    continue
                }
                d := L / 2
                invalid = append(invalid, genInvalidIDs(L, d)...)
            } else {
                for d := 1; d <= L/2; d++ {
                    if L%d != 0 {
                        continue
                    }
                    invalid = append(invalid, genInvalidIDs(L, d)...)
                }
            }
        }

        invalid = remDup(invalid)
        var rangeSum int64
        for _, id := range invalid {
            if id >= start && id <= end {
                rangeSum += id
            }
        }
        total += rangeSum
    }

    return total
}

func main() {

	content := readFile("input.txt")
	sumPart1 := solve(false, content)
	sumPart2 := solve(true, content)

	fmt.Println("Part1:", sumPart1)
	fmt.Println("Part2:", sumPart2)
}