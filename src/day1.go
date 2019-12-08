package adventofcode

import (
	"fmt"
	"log"
	"math"
	"strconv"

	lib "github.com/c-stoeckl/advent-of-code-2019/lib/parseinput"
)

func calcFuel(mass int) int {
	fuel := math.Floor(float64(mass/3)) - 2
	return int(fuel)
}

func day1() {
	lines, err := lib.ParseInput(1)
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
