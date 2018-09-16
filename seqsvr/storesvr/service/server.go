package service

type Server struct {
	ID uint64

	Users  map[uint64]User
	Groups map[uint64]Group
}

func getReq(c *gin.Context) {

}
