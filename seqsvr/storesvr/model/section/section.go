package section

import (
	"fmt"
)

type Section struct {
	ID                 uint64
	GroupLeftInterval  int64
	GroupRightInterval int64
	IsBind             bool
	IP                 string
	BindDate           int64
}

func (s Section) String() string {
	if s.IsBind {
		return fmt.Sprintf("[%d] section interval: [%d:%d] was bind with IP: %s on %d\n",
			s.ID, s.GroupLeftInterval, s.GroupRightInterval, s.IP, s.BindDate)
	}
	return fmt.Sprintf("[%d] Section interval: [%d:%d] was bind",
		s.ID, s.GroupLeftInterval, s.GroupRightInterval)
}
