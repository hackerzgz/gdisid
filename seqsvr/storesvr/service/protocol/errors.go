package protocol

var (
	ErrInvalidReq = Result{
		Code: 101,
		Hint: "invalid request",
	}
	ErrInvalidParam = Result{
		Code: 102,
		Hint: "invalid request params",
	}
)
