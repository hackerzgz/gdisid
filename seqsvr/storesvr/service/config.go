package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackez/gdisid/seqsvr/storesvr/service/protocol"
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
	log.Printf("got GetAllocConfig request: %v\n", *req)

	err = req.Validate()
	if err != nil {
		c.JSON(http.StatusOK, protocol.GetAllocConfigResp{
			Result: protocol.ErrInvalidParam,
		})
		return
	}

	return
}
