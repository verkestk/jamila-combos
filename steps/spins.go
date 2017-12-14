package steps

var Spins = &Family{
	Name: "Spins",
	JL1Steps: []*Step{
		&Step{Name: "Three-Step Turn (Full Spin)", MinCounts: 4},
		&Step{Name: "Open Spin", MinCounts: 2},
		&Step{Name: "4/4 Spin", MinCounts: 4},
		&Step{Name: "2/4 Spin", MinCounts: 2},
	},
	JL2Steps: []*Step{
		&Step{Name: "Diagonal 2/4 Spin", MinCounts: 2},
		&Step{Name: "Centrifugal Spin", MinCounts: 2},
		&Step{Name: "Out-Up-Out-Down Spin", MinCounts: 2},
		&Step{Name: "In-Out Spin", MinCounts: 2},
	},
}
