package adventofcode

import (
	"log"
	"math"
	"strconv"
)

var masses = parseInput(1)
var fuelReqs []int

func calcFuel(mass int) int {
	fuel := math.Floor(float64(mass/3)) - 2
	return int(fuel)
}

func totalFuelReqs() int {
	totalFuelReq := 0

	for _, fuelReq := range fuelReqs {
		totalFuelReq += fuelReq
	}

	return totalFuelReq
}

func day1() {
	for _, mass := range masses {
		mass, err := strconv.Atoi(mass)
		if err != nil {
			log.Fatal(err)
		}

		fuelReq := calcFuel(mass)
		fuelReqs = append(fuelReqs, fuelReq)
	}

	printSolution(part1(), part2())
}

func part1() int {
	return totalFuelReqs()
}

func part2() int {
	for i, fuelReq := range fuelReqs {
		totalModuleFuelReqs := 0

		for ; calcFuel(fuelReq) >= 0; fuelReq = calcFuel(fuelReq) {
			totalModuleFuelReqs += fuelReq
		}

		fuelReqs[i] = totalModuleFuelReqs + fuelReq
	}

	return totalFuelReqs()
}
