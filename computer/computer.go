package computer

import (
	"fmt"
	"log"
)

// Computer is a Intcode computer
type computer struct {
	memory   []int
	index    int
	stepSize int
}

// New creates a new computer instance with defaults
func New() computer {
	c := computer{}
	c.index = 0
	c.stepSize = 4
	return c
}

// Load puts instructions into the computers memory
func (c *computer) Load(program Program) {
	c.memory = program
}

func (c *computer) Execute() []int {
	for {
		c.checkIfIndexOutOfBounds(c.index)
		opcode := OpCode(c.memory[c.index])

		if c.index >= len(c.memory) || opcode == 99 {
			break
		}

		c.interpret(opcode)
		c.index += c.stepSize
	}

	return c.memory
}

func (c *computer) interpret(o OpCode) {
	switch o {
	case 1:
		c.add()
	case 2:
		c.multiply()
	}
}

func (c *computer) add() {
	b := c.nextBlock()
	c.memory[b.outputPos] = c.memory[b.input1Pos] + c.memory[b.input2Pos]
}

func (c *computer) multiply() {
	b := c.nextBlock()
	c.memory[b.outputPos] = c.memory[b.input1Pos] * c.memory[b.input2Pos]
}

func (c *computer) nextBlock() Instruction {
	instruction := Instruction{
		opCode:    c.memory[c.index],
		input1Pos: c.memory[c.index+1],
		input2Pos: c.memory[c.index+2],
		outputPos: c.memory[c.index+3],
	}

	c.checkIfInstructionFaulty(instruction)
	return instruction
}

func (c *computer) checkIfInstructionFaulty(i Instruction) {
	var indexes = [3]int{i.input1Pos, i.input2Pos, i.outputPos}
	for _, index := range indexes {
		c.checkIfIndexOutOfBounds(index)
	}
}

func (c *computer) checkIfIndexOutOfBounds(i int) {
	if i >= len(c.memory) {
		fmt.Printf("Current memory: %v\n", c.memory)
		log.Fatalf("Cannot access memory at index %v", i)
	}
}
