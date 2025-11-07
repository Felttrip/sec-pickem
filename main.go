package main

import (
	"fmt"

	"github.com/cespare/xxhash"
)

type matchup struct {
	TeamA string
	TeamB string
}

var matchups = []matchup{
	{"Georgia", "Mississippi State"},
	{"The Citadel", "Ole Miss"},
	{"Texas A&M", "Missouri"},
	{"Auburn", "Vanderbilt"},
	{"Florida", "Kentucky"},
	{"LSU", "Alabama"},
}

const maxScore = 74

func main() {
	for _, m := range matchups {
		TeamAHash := xxhash.Sum64String(m.TeamA)
		TeamBHash := xxhash.Sum64String(m.TeamB)
		TeamAScore := TeamAHash % maxScore
		TeamBScore := TeamBHash % maxScore
		fmt.Printf("%s vs %s: %d - %d\n", m.TeamA, m.TeamB, TeamAScore, TeamBScore)
		if TeamAScore > TeamBScore {
			fmt.Println(m.TeamA, "wins!\n")
		} else if TeamAScore < TeamBScore {
			fmt.Println(m.TeamB, "wins!\n")
		} else {
			fmt.Println("It's a tie!\n")
		}

	}
}
