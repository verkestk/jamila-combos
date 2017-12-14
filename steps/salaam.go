package steps

var Salaam = &Family{
	Name: "Salaam",
	JL2Steps: []*Step{
		&Step{Name: "Greeting Step", MinCounts: 8},
		&Step{Name: "Salaam Step", MinCounts: 4},
		&Step{Name: "Brush Step", MinCounts: 2},
		&Step{Name: "Bounce Step", MinCounts: 2},
		&Step{Name: "Salaam Step in Circle", MinCounts: 4},
		&Step{Name: "Horse Step", MinCounts: 4},
	},
}
