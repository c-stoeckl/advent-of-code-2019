package adventofcode

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func calcFuel(mass int) int {
	fuel := math.Floor(float64(mass/3)) - 2
	return int(fuel)
}

func part1() {
	lines, err := parseInput(1)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var fuelReqs []int
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		fuel := calcFuel(mass)
		fuelReqs = append(fuelReqs, fuel)
	}

	totalFuelReq := 0
	for _, value := range fuelReqs {
		totalFuelReq += value
	}

	fmt.Print(totalFuelReq)
}

func part2() {

}
