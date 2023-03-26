package graph

import (
	"container/list"
	"fmt"
	"github.com/mbd98/go-vote/v1/lib/primitives"
	"github.com/mbd98/go-vote/v1/lib/util"
)

const (
	MARKER_NONE = 0
	MARKER_TEMP = 1
	MARKER_PERM = 2
)

func (g Graph) FindSources() []primitives.Alternative {
	sources := list.New()
vIter:
	for _, v := range g.Vertices {
		for _, w := range g.Vertices {
			if v != w && g.Edges[w][v] != 0 {
				continue vIter
			}
		}
		sources.PushBack(v)
	}
	return util.ListToArray[primitives.Alternative](sources)
}

func (g Graph) tsVisit(v primitives.Alternative, marker map[primitives.Alternative]uint8, sorted *list.List) error {
	switch marker[v] {
	case MARKER_NONE:
		marker[v] = MARKER_TEMP
		for _, w := range g.Vertices {
			if g.Edges[v][w] != 0 {
				if err := g.tsVisit(w, marker, sorted); err != nil {
					return err
				}
			}
		}
		marker[v] = MARKER_PERM
		sorted.PushFront(v)
		return nil
	case MARKER_TEMP:
		return fmt.Errorf("TopSort: Graph contains cycle")
	case MARKER_PERM:
		return nil
	default:
		return fmt.Errorf("TopSort: Illegal marker %d", marker[v])
	}
}

func (g Graph) TopSort() ([]primitives.Alternative, error) {
	sources := g.FindSources()
	sorted := list.New()
	marker := make(map[primitives.Alternative]uint8, len(g.Vertices))
	for _, v := range sources {
		if err := g.tsVisit(v, marker, sorted); err != nil {
			return nil, err
		}
	}
	for _, v := range g.Vertices {
		if marker[v] != MARKER_PERM {
			if err := g.tsVisit(v, marker, sorted); err != nil {
				return nil, err
			}
		}
	}
	return util.ListToArray[primitives.Alternative](sorted), nil
}