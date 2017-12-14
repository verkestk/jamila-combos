package steps

var HeadMovements = &Family{
	Name: "Head Movements",
	JL1Steps: []*Step{
		&Step{Name: "Head Forward/Back", MinCounts: 2},
		&Step{Name: "Head Side/Side", MinCounts: 2},
		&Step{Name: "Head Around", MinCounts: 2},
		&Step{Name: "Head Crescent", MinCounts: 2},
	},
}
