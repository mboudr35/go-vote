package methods

import (
	"fmt"
	"github.com/mbd98/go-vote/v1/lib/graph"
	"github.com/mbd98/go-vote/v1/lib/primitives"
	"github.com/mbd98/go-vote/v1/lib/util"
	"sort"
	"strings"
)

type rankedAltPair struct {
	primitives.AltPair
	weight int
}

func RankedPairs(g graph.Graph, margin bool) (graph.Graph, string) {
	pairs := make([]rankedAltPair, 0, util.Binomial(len(g.Vertices), 2))
	for _, a := range g.Vertices {
		for _, b := range g.Vertices {
			if a != b && g.Edges[a][b] > g.Edges[b][a] {
				weight := g.Edges[a][b]
				if margin {
					weight -= g.Edges[b][a]
				}
				pairs = append(pairs, rankedAltPair{
					AltPair: primitives.AltPair{
						A: a,
						B: b,
					},
					weight: weight,
				})
			}
		}
	}
	less := func(i, j int) bool {
		return pairs[i].weight > pairs[j].weight
	}
	sort.Slice(pairs, less)
	domg := graph.NewEmptyGraph(g.Vertices)
	for _, pair := range pairs {
		domg.Edges[pair.A][pair.B] = pair.weight
		if domg.CanReach(pair.A, pair.B) {
			// We have a cycle, skip this one
			domg.Edges[pair.A][pair.B] = 0
			pair.weight = 0
		}
	}
	sort.Slice(pairs, less)
	var ineq strings.Builder
	insert := func(pair rankedAltPair) {
		ineq.WriteByte('(')
		ineq.WriteString(fmt.Sprint(pair.A))
		ineq.WriteString(" > ")
		ineq.WriteString(fmt.Sprint(pair.B))
		ineq.WriteByte(')')
	}
	for _, pair := range pairs[0 : len(pairs)-1] {
		if pair.weight != 0 {
			insert(pair)
			ineq.WriteString(" ∧ ")
		}
	}
	if pairs[len(pairs)-1].weight != 0 {
		insert(pairs[len(pairs)-1])
	}

	return domg, ineq.String()
	// TODO: combine the inequalities - maybe use an expression library?
}
