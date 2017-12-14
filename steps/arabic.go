package steps

var Arabic = &Family{
	Name: "Arabic",
	JL1Steps: []*Step{
		&Step{Name: "Arabic 1", MinCounts: 1},
		&Step{Name: "Arabic 2", MinCounts: 2},
		&Step{Name: "Arabic 3", MinCounts: 2},
		&Step{Name: "Arabic 4", MinCounts: 2},
	},
	JL2Steps: []*Step{
		&Step{Name: "Eight Count", MinCounts: 4},
	},
}
