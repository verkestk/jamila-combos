package steps

type Step struct {
	Name      string
	MinCounts int
}

type Family struct {
	Name     string
	JL1Steps []*Step
	JL2Steps []*Step
}

var families = []*Family{
	Arabic,
	BasicEgyptian,
	Debke,
	HeadMovements,
	RunningChooChoo,
	Salaam,
	Shimmy,
	Spins,
	Taqsim,
}

func GetStepsForCounts(counts int, includeLevel2 bool) []*Step {
	steps := []*Step{}
	for _, family := range families {
		familySteps := family.JL1Steps
		if includeLevel2 {
			familySteps = append(familySteps, family.JL2Steps...)
		}

		for _, step := range familySteps {
			if step.MinCounts <= counts && counts%step.MinCounts == 0 {
				numberOfTimes := counts / step.MinCounts
				if numberOfTimes == 1 || numberOfTimes == 2 || numberOfTimes == 4 {
					steps = append(steps, step)
				}
			}
		}
	}

	return steps
}
