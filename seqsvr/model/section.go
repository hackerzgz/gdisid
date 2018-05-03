package model

import (
	"fmt"
	"time"
)

type Section struct {
	ID       uint64
	Number   int32
	IP       uint32
	BindDate int64
}

func (s Section) String() string {
	return fmt.Sprintf("[%d] Section number: %s was bind with IP: %s on %d\n",
		s.ID, s.Number, s.IP, s.BindDate)
}
