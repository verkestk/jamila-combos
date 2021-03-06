package steps

var BasicEgyptian = &Family{
	Name: "Basic Egyptian",
	JL1Steps: []*Step{
		&Step{Name: "Basic Egyptian", MinCounts: 2},
		&Step{Name: "Basic Egyptian Backwalk", MinCounts: 2},
		&Step{Name: "Basic Egyptian Walk with Pivot", MinCounts: 4},
		&Step{Name: "Basic Egyptian Walk with Pivot Angled", MinCounts: 4},
		&Step{Name: "Bow Step", MinCounts: 4},
		&Step{Name: "Step Forward-Back-Forward", MinCounts: 4},
		&Step{Name: "Full Spin with Basic Egyptian", MinCounts: 4},
		&Step{Name: "Open Spin with Basic Egyptian", MinCounts: 8},
		&Step{Name: "Pivot Shift Step", MinCounts: 2},
		&Step{Name: "Half Turn with Pivot Shift Step", MinCounts: 4},
		&Step{Name: "Twist Step", MinCounts: 2},
		&Step{Name: "Twist Tep with Leg Lift", MinCounts: 8},
		&Step{Name: "CCW Pivot Halftime", MinCounts: 1},
		&Step{Name: "CCW Pivot Fulltime", MinCounts: 1},
		&Step{Name: "CCW Pivot Doubletime", MinCounts: 1},
		&Step{Name: "CCW Pivot: One Up, One Down", MinCounts: 2},
		&Step{Name: "CCW Pivot with Leg Lift", MinCounts: 8},
		&Step{Name: "Five Count", MinCounts: 4},
		&Step{Name: "Five Count with Half Spin", MinCounts: 4},
		&Step{Name: "Five Count with Full Spin", MinCounts: 4},
	},
	JL2Steps: []*Step{
		&Step{Name: "Syncopated Pivot Shift Step", MinCounts: 2},
		&Step{Name: "Full Spin with Pivot Shift Step", MinCounts: 4},
		&Step{Name: "Open Spin with Pivot Shift Step", MinCounts: 8},
		&Step{Name: "V Step", MinCounts: 8},
		&Step{Name: "V Step with Spin", MinCounts: 8},
		&Step{Name: "V Step wtih Pivot Shift Step", MinCounts: 8},
		&Step{Name: "V Step with Pitot Shift Step & Spin", MinCounts: 8},
		&Step{Name: "Stomp Step", MinCounts: 2},
		&Step{Name: "CCW Pivot Spin", MinCounts: 2},
		&Step{Name: "Whip Spin with Twist", MinCounts: 8},
		&Step{Name: "Four Count", MinCounts: 2},
		&Step{Name: "X Step", MinCounts: 1},
	},
}
