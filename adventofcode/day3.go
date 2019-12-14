package adventofcode

import (
	"log"
	"math"
	"strconv"
	"strings"
)

// A Path is a slice of coordinates
type Path []Coordinate

// A Coordinate consists is a pair of X and Y positions
type Coordinate struct {
	X int
	Y int
}

// A Vector has a direction and a distance
type Vector struct {
	Direction string
	Distance  int
}

var wires = parseInput(3)

func day3() {
	part1 := d3p1()
	printSolution(part1)
}

func d3p1() int {
	var wirePaths []Path

	for _, wire := range wires {
		wirePath := Path{}
		vectors := strings.Split(wire, ",")
		lastX := 0
		lastY := 0

		for _, v := range vectors {
			direction := string(v[0])
			distance, err := strconv.Atoi(v[1:])
			if err != nil {
				log.Fatalf("Error converting ASCII to Int: %v", err)
			}
			vector := Vector{Direction: direction, Distance: distance}

			switch vector.Direction {
			case "R":
				for newX := lastX; newX <= vector.Distance+lastX; newX++ {
					wirePath = append(wirePath, Coordinate{newX, lastY})
				}
				lastX = vector.Distance + lastX
			case "L":
				for newX := lastX; newX >= lastX-vector.Distance; newX-- {
					wirePath = append(wirePath, Coordinate{newX, lastY})
				}
				lastX = lastX - vector.Distance
			case "U":
				for newY := lastY; newY <= vector.Distance+lastY; newY++ {
					wirePath = append(wirePath, Coordinate{lastX, newY})
				}
				lastY = vector.Distance + lastY
			case "D":
				for newY := lastY; newY >= lastY-vector.Distance; newY-- {
					wirePath = append(wirePath, Coordinate{lastX, newY})
				}
				lastY = lastY - vector.Distance
			}
		}
		wirePaths = append(wirePaths, wirePath)
	}
	crossings := findCrossings(wirePaths[0], wirePaths[1])

	var distances []int
	for _, c := range crossings {
		distances = append(distances, calcManDist(Coordinate{0, 0}, c))
	}

	return findSmallestNumber(distances)
}

func findSmallestNumber(numbers []int) int {
	smallestNumber := numbers[0]
	for _, n := range numbers {
		if n < smallestNumber {
			smallestNumber = n
		}
	}

	return smallestNumber
}

func calcManDist(a, b Coordinate) int {
	return int(math.Abs(float64(a.X)-float64(b.X))) + int(math.Abs(float64(a.Y)-float64(b.Y)))
}

func findCrossings(path1, path2 Path) (crossings []Coordinate) {
	m := make(map[Coordinate]bool)

	for _, c := range path1 {
		m[c] = true
	}

	for _, c := range path2 {
		if _, ok := m[c]; ok {
			if c.X == 0 && c.Y == 0 {
				continue
			}
			crossings = append(crossings, c)
		}
	}

	return crossings
}
