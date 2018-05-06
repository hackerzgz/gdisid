package section

import (
	"fmt"
)

type Section struct {
	ID       uint64
	Number   int32
	IP       string
	BindDate int64
}

func (s Section) String() string {
	return fmt.Sprintf("[%d] Section number: %d was bind with IP: %s on %d\n",
		s.ID, s.Number, s.IP, s.BindDate)
}
