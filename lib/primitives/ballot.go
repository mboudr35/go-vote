package primitives

import (
	"fmt"
	"strings"
)

type PreferentialBallot map[Alternative]int

func (b PreferentialBallot) RankingsArray() [][]Alternative {
	rankings := make([][]Alternative, len(b))
	for a, r := range b {
		rankings[r-1] = append(rankings[r-1], a)
	}
	trim := 0
	for i := len(rankings); i >= 0 && rankings[i] == nil; i++ {
		trim++
	}
	return rankings[0 : len(rankings)-trim]
}

func (b PreferentialBallot) String() string {
	rankings := b.RankingsArray()
	var sb strings.Builder
	var i, j int
	var ranking []Alternative
	for i = 0; i < len(rankings)-1; i++ {
		ranking = rankings[i]
		for j = 0; j < len(ranking)-1; j++ {
			sb.WriteString(fmt.Sprint(ranking[j]))
			sb.WriteString(" = ")
		}
		sb.WriteString(fmt.Sprint(ranking[j]))
		sb.WriteString(" > ")
	}
	ranking = rankings[i]
	for j = 0; j < len(ranking)-1; j++ {
		sb.WriteString(fmt.Sprint(ranking[j]))
		sb.WriteString(" = ")
	}
	sb.WriteString(fmt.Sprint(ranking[j]))
	return sb.String()
}
