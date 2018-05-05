package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hackez/gdisid/seqsvr/storesvr/service"
)

var (
	r *gin.Engine
)

// Register http router support server for alloc server
func Register() {
	r = gin.Default()
	basic := r.Group("api/")
	{
		alloc := basic.Group("/alloc")

		config := alloc.Group("/config")
		config.GET("/get", service.GetAllocConfig) // TODO: get alloc start config

		token := alloc.Group("/token")
		token.POST("/refresh", service.RefreshToken) // TODO: refresh alloc server token

		uid := alloc.Group("/uid")
		uid.POST("/update", service.UpdateUID)
	}
}

// Run start store server under address
// listen on port :8080 if address is empty
func Run(addr ...string) error {
	return r.Run(addr...)
}
