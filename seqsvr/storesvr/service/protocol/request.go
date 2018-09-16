package protocol

import (
	"errors"
)

// Request send message for any interafce
type Request interface {
	Validate() error
}

// GetAllocConfigReq request for GetAllocConfig infterface
type GetAllocConfigReq struct {
	IP string `json:"ip"`
}

func (req *GetAllocConfigReq) Validate() error {
	if req.IP == "" {
		return errors.New("empty ip")
	}
	return nil
}
