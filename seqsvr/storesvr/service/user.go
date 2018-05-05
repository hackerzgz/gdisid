package service

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint64
	Name     string
	Sequence uint64

	GroupID uint64
}

// GetSeq return next non-repeating sequence for user
func (u *User) GetSeq() uint64 {
	// TODO: compared max sequence in group
	return u.Sequence
}

func UpdateUID(c *gin.Context) {

}
