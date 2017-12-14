package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/verkestk/jamila-combos/steps"
)

const (
	// set this to false to use only JL1 steps
	includeJL2 = true
)

var (
	// code randomly selects one of the following as total count number
	// must be a factor of 8
	counts = []int{
		8,
		16,
		32,
	}

	// code randomly uses these divisions of 8 counts when building combos
	// each must add up to a total of 8
	countsOf8 = [][]int{
		[]int{8},
		[]int{6, 2},
		[]int{4, 4},
		[]int{4, 2, 2},
		[]int{2, 6},
		[]int{2, 4, 2},
		[]int{2, 2, 4},
		[]int{2, 2, 2, 2},
	}
)

func main() {

	rand.Seed(time.Now().Unix())

	counts := selectCounts()
	previousStepName := ""
	totalCounts := 0

	for _, count := range counts {
		possibleSteps := steps.GetStepsForCounts(count, includeJL2)
		step := possibleSteps[rand.Intn(len(possibleSteps))]

		for step.Name == previousStepName {
			step = possibleSteps[rand.Intn(len(possibleSteps))]
		}

		totalCounts += count
		fmt.Printf("%d counts of %s\n", count, step.Name)

		previousStepName = step.Name
	}

	fmt.Printf("\n%d total counts\n", totalCounts)
}

func selectCounts() []int {
	totalCounts := counts[rand.Intn(len(counts))]
	countsSoFar := 0

	counts := []int{}
	for countsSoFar < totalCounts {
		countsSoFar += 8
		counts = append(counts, countsOf8[rand.Intn(len(countsOf8))]...)
	}

	return counts
}
