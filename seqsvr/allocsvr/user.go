package allocsvr

type User struct {
	ID uint64

	CurSeq uint64

	GroupID uint64 // to find max sequence
}
