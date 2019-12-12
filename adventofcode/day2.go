package adventofcode

import (
	"strings"

	c "github.com/c-stoeckl/advent-of-code-2019/computer"
)

func day2() {
	part1 := p1()
	part2 := p2()
	printSolution(part1, part2)
}

func p1() int {
	var newProgram = stringArrayToInt(strings.Split(parseInput(2)[0], ","))

	newProgram[1] = 12
	newProgram[2] = 2

	computer := c.New()
	computer.Load(newProgram)
	result := computer.Execute()
	return result[0]
}

func p2() int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			var newProgram = stringArrayToInt(strings.Split(parseInput(2)[0], ","))
			newProgram[1] = noun
			newProgram[2] = verb
			computer := c.New()
			computer.Load(newProgram)
			result := computer.Execute()
			if result[0] == 19690720 {
				return (100 * noun) + verb
			}
		}
	}
	return 0
}
