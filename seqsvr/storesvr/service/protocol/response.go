package protocol

import (
	"fmt"
)

// Response result for any interface
type Response interface {
	Format() string
}

// Result return error code and hint in every response
type Result struct {
	Code int    `json:"code"`
	Hint string `json:"hint"`
}

// GetAllocConfigResp return alloc config if avaliable
type GetAllocConfigResp struct {
	Result             Result `json:"result"`
	GroupLeftInterval  int64  `json:"group_left_interval"`
	GroupRightInterval int64  `json:"group_right_interval"`
}

func (resp GetAllocConfigResp) Format() string {
	return fmt.Sprintf("%+v\n", resp)
}
