package graph

import (
	"container/list"
	"github.com/mbd98/go-vote/lib/v1/primitives"
	"github.com/mbd98/go-vote/lib/v1/util"
)

func (g Graph) CanReach(target, start primitives.Alternative) bool {
	stack := list.New()
	stack.PushFront(start)
	for stack.Len() > 0 {
		v := stack.Remove(stack.Front()).(primitives.Alternative)
		if v == target {
			return true
		}
		for w, e := range g.Edges[v] {
			if e != 0 {
				stack.PushFront(w)
			}
		}
	}
	return false
}

func (g Graph) DominantWidestPaths() Graph {
	p := NewEmptyGraph(g.Vertices)
	for _, a := range g.Vertices {
		for _, b := range g.Vertices {
			if a != b && g.Edges[a][b] > g.Edges[b][a] {
				p.Edges[a][b] = g.Edges[a][b]
			}
		}
	}
	for _, a := range p.Vertices {
		for _, b := range p.Vertices {
			if a != b {
				for _, c := range p.Vertices {
					if a != c && b != c {
						p.Edges[b][c] = util.Max(p.Edges[b][c], util.Min(p.Edges[b][a], p.Edges[a][c]))
					}
				}
			}
		}
	}
	return p
}
