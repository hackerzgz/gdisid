package user

import (
	"fmt"
)

// User really existing users persist corresponding
// personal information in the databases
type User struct {
	ID       uint64
	Name     string
	Sequence uint64
	GroupID  uint64
}

func (u User) String() string {
	return fmt.Sprintf("[gid: %d]user %s[%d] current sequence: %d\n",
		u.GroupID, u.Name, u.ID, u.Sequence)
}
