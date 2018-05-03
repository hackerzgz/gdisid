package allocsvr

type Server struct {
	ID uint64

	SectionNumber []int32
}

func New() (*Server, error) {
	// should apply config from store server
}
