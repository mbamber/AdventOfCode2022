package main

import (
	"context"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var maxScenicScore int
	for y, line := range lines {
		for x, c := range line {

			// No point checking the edges, as their scenic score will always be 0
			if x == 0 || y == 0 || x == len(line)-1 || y == len(lines)-1 {
				continue
			}

			n := 0
			e := 0
			s := 0
			w := 0

			cint, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, err
			}

			// North
			for y1 := y - 1; y1 >= 0; y1-- {
				i, err := strconv.Atoi(string(lines[y1][x]))
				if err != nil {
					return nil, err
				}

				if cint <= i {
					n++
					break
				}
				n++
			}

			// East
			for x1 := x + 1; x1 < len(line); x1++ {
				i, err := strconv.Atoi(string(lines[y][x1]))
				if err != nil {
					return nil, err
				}

				if cint <= i {
					e++
					break
				}
				e++
			}

			// South
			for y2 := y + 1; y2 < len(lines); y2++ {
				i, err := strconv.Atoi(string(lines[y2][x]))
				if err != nil {
					return nil, err
				}

				if cint <= i {
					s++
					break
				}
				s++
			}

			// West
			for x2 := x - 1; x2 >= 0; x2-- {
				i, err := strconv.Atoi(string(lines[y][x2]))
				if err != nil {
					return nil, err
				}

				if cint <= i {
					w++
					break
				}
				w++
			}

			scenicScore := n * e * s * w
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return maxScenicScore, nil
}
