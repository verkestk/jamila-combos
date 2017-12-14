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

/*func selectSteps(number int) []string {
  allPossibleSteps := getPossibleSteps()
  selectedSteps := []string {}

  for len(selectedSteps) < number {
    selectedSteps = append(selectedSteps, selectStep(allPossibleSteps))
  }

  return selectedSteps
}

func getPossibleSteps() []string{
  steps := []string{}

  for _, family := range families {
    jl1Steps, ok := jl1[family]
    if includeJL1 && ok {
      steps = append(steps, jl1Steps...)
    }

    jl2Steps, ok := jl2[family]
    if includeJL2 && ok {
      steps = append(steps, jl2Steps...)
    }
  }

  return steps
}

func selectStep(steps []string) string{
  return steps[rand.Intn(len(steps))]
}*/

/*

families = []string {
  "BE",
  "Arabic",
  "RChooChoo",
  "Shimmy",
  "Salaam",
  "Debke",
  "Taqsim",
  "Spin",
  "Head",
}

jl1 = map[string][]string{
  "BE": []string{
    "Basic Egyptian",
    "Basic Egyptian Backwalk",
    "Basic Egyptian Walk with Pivot",
    "Basic Egyptian Walk with Pivot Angled",
    "Bow Step",
    "Step Forward-Back-Forward",
    "Full Spin with Basic Egyptian",
    "Open Spin with Basic Egyptian",
    "Pivot Shift Step",
    "Half Turn with Pivot Shift Step",
    "Twist Step",
    "Twist Tep with Leg Lift",
    "CCW Pivot Halftime",
    "CCW Pivot Fulltime",
    "CCW Pivot Doubletime",
    "CCW Pivot: One Up, One Down",
    "CCW Pivot with Leg Lift",
    "Five Count",
    "Five Count with Half Spin",
    "Five Count with Full Spin",
  },
  "Arabic": []string{
    "Arabic 1",
    "Arabic 2",
    "Arabic 3",
    "Arabic 4",
  },
  "RChooChoo": []string{
    "Running Choo-Choo",
    "4 Forward, 4 Back",
    "2 Forward, 2 Back",
    "1 Forward, 1 Back",
    "Forward, Middle, Middle, Back",
    "Zanouba",
  },
  "Shimmy": []string{
    "Singles on the Up Halftime",
    "Singles on the Up Fulltime",
  },
  "Taqsim": []string{
    "Basic Taqsim",
    "Reverse Basic Taqsim",
    "Maya",
    "Reverse Maya",
    "Circle Step",
    "Crescent Step",
  },
  "Spins": []string{
    "Three-Step Turn (Full Spin)",
    "Open Spin",
    "4/4 Spin",
    "2/4 Spin",
  },
  "Head": []string{
    "Head Forward/Back",
    "Head Side/Side",
    "Head Around",
    "Head Crescent",
  },
}

jl2 = map[string][]string{
  "BE": []string{
    "Syncopated Pivot Shift Step",
    "Full Spin with Pivot Shift Step",
    "Open Spin with Pivot Shift Step",
    "V Step",
    "V Step with Spin",
    "V Step wtih Pivot Shift Step",
    "V Step with Pitot Shift Step & Spin",
    "Stomp Step",
    "CCW Pivot Spin",
    "Whip Spin with Twist",
    "Four Count",
    "X Step",
  },
  "Arabic": []string{
    "Eight Count",
  },
  "Shimmy": []string{
    "Choo-Choo",
    "Shimmy Spin",
    "Stomp Step with Shimmy",
    "Algerian Shimmy",
    "Four-Four Shimmy",
    "Singles on the Down",
    "Ahmad Shimmy",
    "3/4 Shimmy: 3/4 on the Up",
    "F-and-B-and Walk-2-3-4",
    "Three Quarter Shimmy Spin",
    "Three Quarter Shimmy with Twist",
    "One, Two, Three, And",
    "Three Quarter Flamenco",
    "Sahima",
  },
  "Salaam": []string {
    "Greeting Step",
    "Salaam Step",
    "Brush Step",
    "Bounce Step",
    "Salaam Step in Circle",
    "Horse Step",
  },
  "Debke": []string {
    "Debke 1 (basic)",
    "Debke 2 (F & B &)",
    "Debke 3 (chasse swing)",
    "Debke 4 (brush stomp)",
    "Debke 5 (hop heel dig)",
  },
  "Taqsim": []string{
    "Crescent Step with Pelvic Locks",
    "Turkish Walk",
    "Turkish Backwalk",
    "Pyramid Step / Suzi Q",
    "Goosh Step",
    "Goosh Spin",
    "F8 Backwalk",
    "Rib Figure Eights & Chest Locks",
  },
  "Spins": []string{
    "Diagonal 2/4 Spin",
    "Centrifugal Spin",
    "Out-Up-Out-Down Spin",
    "In-Out Spin",
  },
}
*/
