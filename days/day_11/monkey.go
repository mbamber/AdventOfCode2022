package main

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/knetic/govaluate"
)

type Monkey struct {
	items           []int
	postInspectFunc func(m *Monkey) error
	inspectFunc     func(m *Monkey) error
	testDivisibleBy int
	testTrueMonkey  int
	testFalseMonkey int

	inspectionCount int
}

func Parse(s string) (*Monkey, error) {
	m := &Monkey{
		postInspectFunc: func(m *Monkey) error { return nil },
	}

	for _, line := range strings.Split(s, "\n") {
		l := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(l, "Starting items:"):
			itemsString := strings.TrimPrefix(l, "Starting items: ")
			items := strings.Split(itemsString, ", ")
			for _, item := range items {
				i, err := strconv.Atoi(item)
				if err != nil {
					return nil, err
				}
				m.items = append(m.items, i)
			}
		case strings.HasPrefix(l, "Operation:"):
			expr, err := govaluate.NewEvaluableExpression(strings.TrimPrefix(l, "Operation: new = "))
			if err != nil {
				return nil, err
			}

			m.inspectFunc = func(m *Monkey) error {
				res, err := expr.Evaluate(map[string]interface{}{"old": m.GetCurrentWorry()})
				if err != nil {
					return err
				}

				r, ok := res.(float64)
				if !ok {
					return errors.New("response was not a float64")
				}

				m.UpdateCurrentWorry(int(r))
				return nil
			}
		case strings.HasPrefix(l, "Test: divisible by"):
			x, err := strconv.Atoi(strings.TrimPrefix(l, "Test: divisible by "))
			if err != nil {
				return nil, err
			}
			m.testDivisibleBy = x
		case strings.HasPrefix(l, "If true: throw to monkey"):
			x, err := strconv.Atoi(strings.TrimPrefix(l, "If true: throw to monkey "))
			if err != nil {
				return nil, err
			}
			m.testTrueMonkey = x
		case strings.HasPrefix(l, "If false: throw to monkey"):
			x, err := strconv.Atoi(strings.TrimPrefix(l, "If false: throw to monkey "))
			if err != nil {
				return nil, err
			}
			m.testFalseMonkey = x
		}
	}

	return m, nil
}

func (m *Monkey) HasItemsToInspect() bool {
	return len(m.items) > 0
}

// Inspect the next item
// Panic if there are no items to inspect
func (m *Monkey) Inspect() {
	m.inspectionCount++
	m.inspectFunc(m)
	m.postInspectFunc(m)
}

// Throw the next item to the monkey with the given id
// Panic if there no items to inspect
func (m *Monkey) Throw() (item int, monkey int) {
	item = m.items[0]
	m.items = m.items[1:]

	if math.Mod(float64(item), float64(m.testDivisibleBy)) == 0 {
		monkey = m.testTrueMonkey
	} else {
		monkey = m.testFalseMonkey
	}

	return item, monkey
}

// Catch an item
func (m *Monkey) Catch(i int) {
	m.items = append(m.items, i)
}

func (m *Monkey) Activity() int {
	return m.inspectionCount
}

func (m *Monkey) GetCurrentWorry() int {
	return m.items[0]
}

func (m *Monkey) UpdateCurrentWorry(n int) {
	m.items[0] = n
}

func (m *Monkey) SetPostInspect(f func(m *Monkey) error) {
	m.postInspectFunc = f
}

func (m *Monkey) GetTestCondition() int {
	return m.testDivisibleBy
}
