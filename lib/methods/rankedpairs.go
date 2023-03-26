package methods

import (
	"github.com/mbd98/go-vote/v1/lib/graph"
	"github.com/mbd98/go-vote/v1/lib/primitives"
	"github.com/mbd98/go-vote/v1/lib/util"
	"sort"
)

type rankedAltPair struct {
	primitives.AltPair
	weight int
}

func RankedPairs(g graph.Graph, margin bool) (graph.Graph, []primitives.Alternative, error) {
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
	dom := graph.NewEmptyGraph(g.Vertices)
	for _, pair := range pairs {
		dom.Edges[pair.A][pair.B] = pair.weight
		if dom.CanReach(pair.A, pair.B) {
			// We have a cycle, skip this one
			dom.Edges[pair.A][pair.B] = 0
			pair.weight = 0
		}
	}
	sort.Slice(pairs, less)

	ts, err := dom.TopSort()
	if err != nil {
		return graph.Graph{}, nil, err
	}
	return dom, ts, nil
}
