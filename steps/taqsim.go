package steps

var Taqsim = &Family{
	Name: "Taqsim",
	JL1Steps: []*Step{
		&Step{Name: "Basic Taqsim", MinCounts: 2},
		&Step{Name: "Reverse Basic Taqsim", MinCounts: 2},
		&Step{Name: "Maya", MinCounts: 2},
		&Step{Name: "Reverse Maya", MinCounts: 2},
		&Step{Name: "Circle Step", MinCounts: 2},
		&Step{Name: "Crescent Step", MinCounts: 2},
	},
	JL2Steps: []*Step{
		&Step{Name: "Crescent Step with Pelvic Locks", MinCounts: 2},
		&Step{Name: "Turkish Walk", MinCounts: 4},
		&Step{Name: "Turkish Backwalk", MinCounts: 4},
		&Step{Name: "Pyramid Step / Suzi Q", MinCounts: 4},
		&Step{Name: "Goosh Step", MinCounts: 2},
		&Step{Name: "Goosh Spin", MinCounts: 4},
		&Step{Name: "F8 Backwalk", MinCounts: 4},
		&Step{Name: "Rib Figure Eights & Chest Locks", MinCounts: 2},
	},
}
