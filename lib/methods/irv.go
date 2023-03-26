package methods

import (
	"github.com/mbd98/go-vote/v1/lib/primitives"
	"github.com/mbd98/go-vote/v1/lib/util"
	"math"
)

func eliminateAlt(eliminate primitives.Alternative, alts []primitives.Alternative, ballots []primitives.PreferentialBallot) ([]primitives.Alternative, []primitives.PreferentialBallot) {
	newBallots := make([]primitives.PreferentialBallot, len(ballots))
	for i, ballot := range ballots {
		newBallots[i] = make(primitives.PreferentialBallot, len(alts))
		r := ballot[eliminate]
		//delete(ballot, eliminate)
		for alt, rank := range ballot {
			if alt != eliminate {
				if rank > r {
					newBallots[i][alt] = ballot[alt] - 1
				} else {
					newBallots[i][alt] = ballot[alt]
				}
			}
		}
	}
	return util.ArrayDeleteElement(alts, eliminate), newBallots
}

func InstantRunoff(alts []primitives.Alternative, ballots []primitives.PreferentialBallot) primitives.Alternative {
	allocation := make(map[primitives.Alternative]int, len(alts))

	// Count voters' first choices
	for _, ballot := range ballots {
		for alt, rank := range ballot {
			if rank == 1 {
				allocation[alt]++
			}
		}
	}

	var fewestAlt primitives.Alternative
	fewestCount := math.MaxInt

	// Anyone have a majority?
	for alt, votes := range allocation {
		if 2*votes > len(ballots) {
			// We have a winner!
			return alt
		}
		if fewestCount > votes {
			fewestCount = votes
			fewestAlt = alt
		}
	}

	// Eliminate the biggest loser, try again
	return InstantRunoff(eliminateAlt(fewestAlt, alts, ballots))
}
