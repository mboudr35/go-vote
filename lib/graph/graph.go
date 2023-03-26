package graph

import (
	"fmt"
	"github.com/mbd98/go-vote/lib/primitives"
	"strings"
)

type Graph struct {
	Vertices []primitives.Alternative
	Edges    map[primitives.Alternative]map[primitives.Alternative]int
}

func NewEmptyGraph(vertices []primitives.Alternative) Graph {
	edges := make(map[primitives.Alternative]map[primitives.Alternative]int, len(vertices))
	for _, v := range vertices {
		edges[v] = make(map[primitives.Alternative]int, len(vertices))
	}
	return Graph{
		Vertices: vertices,
		Edges:    edges,
	}
}

func NewElectionGraph(alts []primitives.Alternative, ballots []primitives.PreferentialBallot) Graph {
	g := NewEmptyGraph(alts)
	for _, ballot := range ballots {
		for x, xr := range ballot {
			for y, yr := range ballot {
				if x != y && xr < yr {
					g.Edges[x][y]++
				}
			}
		}
	}
	return g
}

func (g Graph) String() string {
	return fmt.Sprint(g.Edges)
}

func (g Graph) PrettyString() string {
	var sb strings.Builder
	sb.WriteByte('\t')
	for _, v := range g.Vertices {
		sb.WriteString(fmt.Sprint(v))
		sb.WriteByte('\t')
	}
	sb.WriteByte('\n')
	for _, src := range g.Vertices {
		sb.WriteString(fmt.Sprint(src))
		sb.WriteByte('\t')
		for _, dst := range g.Vertices {
			sb.WriteString(fmt.Sprint(g.Edges[src][dst]))
			sb.WriteByte('\t')
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
