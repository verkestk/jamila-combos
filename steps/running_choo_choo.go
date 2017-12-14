package steps

var RunningChooChoo = &Family{
	Name: "Running Choo Choo",
	JL1Steps: []*Step{
		&Step{Name: "Running Choo-Choo", MinCounts: 1},
		&Step{Name: "4 Forward, 4 Back", MinCounts: 8},
		&Step{Name: "2 Forward, 2 Back", MinCounts: 4},
		&Step{Name: "1 Forward, 1 Back", MinCounts: 2},
		&Step{Name: "Forward, Middle, Middle, Back", MinCounts: 4},
		&Step{Name: "Zanouba", MinCounts: 2},
	},
}
