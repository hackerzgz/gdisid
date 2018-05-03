package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

// Register http router support server for alloc server
func Register() {
	basic := r.Group("api/")
	{
		alloc := basic.Group("/alloc")
		alloc.GET("/config/get") // TODO: get alloc start config
	}
}

// Run start store server under address
// listen on port :8080 if address is empty
func Run(addr ...string) error {
	return r.Run(addr...)
}
