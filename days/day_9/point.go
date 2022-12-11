package main

import (
	"math"
)

type Point struct {
	X, Y int
}

func (p *Point) IsAdjacent(p2 *Point) bool {
	xDiff := math.Abs(float64(p.X)-float64(p2.X)) == 1.0
	yDiff := math.Abs(float64(p.Y)-float64(p2.Y)) == 1.0

	return (p.X == p2.X && p.Y == p2.Y) || (p.X == p2.X && yDiff) || (p.Y == p2.Y && xDiff) || (xDiff && yDiff)
}

func (p *Point) MoveTowards(p2 *Point) {
	if p.X == p2.X {
		if p.Y > p2.Y {
			p.Y--
		} else if p.Y < p2.Y {
			p.Y++
		}
	} else if p.Y == p2.Y {
		if p.X > p2.X {
			p.X--
		} else if p.X < p2.X {
			p.X++
		}
	} else {
		// Move diagonally
		xPositive := (p2.X - p.X) > 0
		yPositive := (p2.Y - p.Y) > 0

		if xPositive {
			p.X++
		} else {
			p.X--
		}

		if yPositive {
			p.Y++
		} else {
			p.Y--
		}
	}
}

func (p *Point) Copy() *Point {
	return &Point{p.X, p.Y}
}

type Locations []*Point

func (l Locations) CountUnique() int {
	var unique int

i:
	for i, p := range l {
		for i2 := i + 1; i2 < len(l); i2++ {
			if p.X == l[i2].X && p.Y == l[i2].Y {
				continue i
			}
		}
		unique++
	}

	return unique
}
