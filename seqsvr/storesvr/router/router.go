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

	// register alloc group api
	registerAlloc(basic)
}

// registerAlloc api group for alloc request
func registerAlloc(group *gin.RouterGroup) {
	allocG := group.Group("/alloc")

	// config group provide alloc service running config
	confG := allocG.Group("/config")
	confG.GET("/get", service.GetAllocConfig) // TODO: get alloc start config

	// token group provide alloc service validity check
	tokenG := allocG.Group("/token")
	tokenG.POST("/refresh", service.RefreshToken) // TODO: refresh alloc server token

	// uid group provide newest distribution uid for alloc service
	uidG := allocG.Group("/uid")
	uidG.POST("/update", service.UpdateUID) // TODO(hackerzgz): update uid steps and return newest uid
}

// Run start store server under address
// listen on port :8080 if address is empty
func Run(addr ...string) error {
	return r.Run(addr...)
}
