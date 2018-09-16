package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackez/gdisid/seqsvr/storesvr/service/protocol"
	"github.com/hackez/zapwrapper"
	"go.uber.org/zap"
)

// GetAllocConfig return avaliable config
// to running alloc service
func GetAllocConfig(c *gin.Context) {
	req := new(protocol.GetAllocConfigReq)
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, protocol.GetAllocConfigResp{
			Result: protocol.ErrInvalidReq,
		})
		return
	}
	zapwrapper.Logger.Info(
		"got GatAllocConfig request:",
		zap.String("ip:", req.IP))

	err = req.Validate()
	if err != nil {
		c.JSON(http.StatusOK, protocol.GetAllocConfigResp{
			Result: protocol.ErrInvalidParam,
		})
		return
	}

	return
}
