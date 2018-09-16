package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackez/gdisid/seqsvr/storesvr/service/protocol"
	"github.com/hackez/zapwrapper"
	"go.uber.org/zap"
)

type Server struct {
	ID uint64

	Users  map[uint64]User
	Groups map[uint64]Group
}

// getReq parse and valid request
func getReq(c *gin.Context, req protocol.Request) {
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, protocol.GetAllocConfigResp{
			Result: protocol.ErrInvalidReq,
		})
		return
	}

	err = req.Validate()
	if err != nil {
		c.JSON(http.StatusOK, protocol.GetAllocConfigResp{
			Result: protocol.ErrInvalidParam,
		})
		return
	}
}
