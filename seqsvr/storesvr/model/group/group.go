package group

import (
	"fmt"
)

// Group user id adjacent users to combine a group
type Group struct {
	ID          uint64
	MaxSequence uint64
}

func (g *Group) String() string {
	return fmt.Sprintf("group id: %d with max sequence: %d",
		g.ID, g.MaxSequence)
}
