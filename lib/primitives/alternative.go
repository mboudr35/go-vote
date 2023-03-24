package primitives

import (
	"github.com/google/uuid"
)

type Alternative struct {
	Id   uuid.UUID
	Name string
}

func (alt Alternative) String() string {
	return alt.Name
}

type AltPair struct {
	A, B Alternative
}
