package methods

import (
	"github.com/mbd98/go-vote/lib/v1/graph"
	"github.com/mbd98/go-vote/lib/v1/primitives"
)

func Schulze(g graph.Graph, margin bool) (graph.Graph, map[primitives.Alternative]bool, map[primitives.Alternative]map[primitives.Alternative]bool) {
	d := g
	if margin {
		d = graph.NewEmptyGraph(g.Vertices)
		for _, a := range g.Vertices {
			for _, b := range g.Vertices {
				d.Edges[a][b] = g.Edges[a][b] - g.Edges[b][a]
			}
		}
	}
	p := d.DominantWidestPaths()
	winner := make(map[primitives.Alternative]bool, len(p.Vertices))
	dom := make(map[primitives.Alternative]map[primitives.Alternative]bool, len(g.Vertices))
	for _, a := range p.Vertices {
		dom[a] = make(map[primitives.Alternative]bool, len(g.Vertices))
		winner[a] = true
		for _, b := range p.Vertices {
			if a != b {
				if p.Edges[b][a] > p.Edges[a][b] {
					dom[a][b] = false
					winner[a] = false
				} else {
					dom[a][b] = true
				}
			}
		}
	}
	return p, winner, dom
}
