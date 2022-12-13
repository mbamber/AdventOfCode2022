package main

import (
	"errors"
	"strconv"
	"strings"
)

type Status string

const (
	StatusWaiting   Status = "WAITING"
	StatusExecuting Status = "EXECUTING"
	StatusComplete  Status = "COMPLETE"
)

func ParseInstruction(s string) (Instruction, error) {
	fields := strings.Fields(s)
	if len(fields) <= 0 {
		return nil, errors.New("invalid instruction")
	}

	switch fields[0] {
	case "noop":
		return NewInstructionNoop(), nil
	case "addx":
		if len(fields) != 2 {
			return nil, errors.New("no argument to addx")
		}
		x, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		return NewInstructionAddX(x), nil
	default:
		return nil, errors.New("unknown instruction")
	}
}

type Instruction interface {
	Start(c *CPU) error
	Continue(c *CPU) error
	Complete(c *CPU) error

	Status() Status
}

type InstructionNoop struct {
	status Status
	cycles int
}

func NewInstructionNoop() *InstructionNoop {
	return &InstructionNoop{
		status: StatusWaiting,
		cycles: 1,
	}
}

func (i *InstructionNoop) Start(c *CPU) error {
	i.status = StatusExecuting
	return nil
}

func (i *InstructionNoop) Continue(c *CPU) error {
	i.cycles--
	if i.cycles == 0 {
		i.status = StatusComplete
	}
	return nil
}

func (i *InstructionNoop) Complete(c *CPU) error {
	return nil
}

func (i *InstructionNoop) Status() Status {
	return i.status
}

type InstructionAddX struct {
	status Status
	cycles int
	x      int
}

func NewInstructionAddX(x int) *InstructionAddX {
	return &InstructionAddX{
		status: StatusWaiting,
		cycles: 2,
		x:      x,
	}
}

func (i *InstructionAddX) Start(c *CPU) error {
	i.status = StatusExecuting
	return nil
}

func (i *InstructionAddX) Continue(c *CPU) error {
	i.cycles--
	if i.cycles == 0 {
		i.status = StatusComplete
	}
	return nil
}

func (i *InstructionAddX) Complete(c *CPU) error {
	c.x += i.x
	return nil
}

func (i *InstructionAddX) Status() Status {
	return i.status
}

type Hook func(c *CPU) error

// A Clocked CPU
type CPU struct {
	ip      int
	program []Instruction

	cycle int // The current cycle

	hooks map[Status][]Hook // Inspection hooks that run during each cycle

	// Registers
	x int
}

type Opt func(c *CPU)

func WithHook(s Status, h Hook) Opt {
	return func(c *CPU) {
		if c.hooks[s] == nil {
			c.hooks[s] = []Hook{h}
		} else {
			c.hooks[s] = append(c.hooks[s], h)
		}
	}
}

func NewCPU(program []Instruction, opts ...Opt) *CPU {
	c := &CPU{
		ip:      0,
		program: program,

		cycle: 0,

		hooks: map[Status][]Hook{},

		x: 1,
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *CPU) HasMoreInstructions() bool {
	return c.ip < len(c.program)
}

func (c *CPU) Tick() error {
	c.cycle++
	i := c.program[c.ip]

	if i.Status() == StatusWaiting {
		err := c.runHooks()
		if err != nil {
			return err
		}

		err = i.Start(c)
		if err != nil {
			return err
		}
	}

	if i.Status() == StatusExecuting {
		err := c.runHooks()
		if err != nil {
			return err
		}

		err = i.Continue(c)
		if err != nil {
			return err
		}
	}

	if i.Status() == StatusComplete {
		err := c.runHooks()
		if err != nil {
			return err
		}

		err = i.Complete(c)
		if err != nil {
			return err
		}

		c.ip++
	}

	return nil
}

func (c *CPU) runHooks() error {
	hooks, ok := c.hooks[c.program[c.ip].Status()]
	if ok {
		for _, h := range hooks {
			err := h(c)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *CPU) CurrentCycle() int {
	return c.cycle
}

func (c *CPU) RegisterX() int {
	return c.x
}
