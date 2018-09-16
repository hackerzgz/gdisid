package service

import (
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackez/gdisid/seqsvr/common/inet"
	msection "github.com/hackez/gdisid/seqsvr/storesvr/model/section"
	"github.com/hackez/gdisid/seqsvr/storesvr/service/protocol"
	"github.com/hackez/zapwrapper"
	"go.uber.org/zap"
)

// GetAllocConfig return avaliable config
// to running alloc service
func GetAllocConfig(c *gin.Context) {
	req := new(protocol.GetAllocConfigReq)
	getReq(c, req)
	zapwrapper.Logger.Info(
		"got GatAllocConfig request:",
		zap.String("ip:", req.IP))

	// invoke vacant user group
	sections, err := msection.Get(
		"is_bind",
		"left_interval DESC", 0, 1, false)
	if err != nil {
		zapwrapper.Logger.Error(
			"false to get section",
			zap.String("error", err.Error()),
		)

		// TODO(hackerzgz): return common error message
		return
	}

	if len(sections) == 0 { // no avaliable section
		// TODO(hackerzgz): return no avaliable error message
		return
	}

	// binding this machine to section
	affected, err := msection.Modify(msection.Section{
		ID:       sections[0].ID,
		IP:       req.IP,
		IsBind:   false,
		BindDate: time.Now().Unix(),
	})
	if err != nil {
		// TODO(hackerzgz): return common error message
		return
	}

	// TODO(hackerzgz): register ip to vacant user group
	return
}

func UpdateSection(c *gin.Context) {
	req := new(protocol.UpdateSectionReq)
	getReq(c, req)
	zapwrapper.Logger.Info(
		"got UpdateSection request:",
		zap.String("left_interval:", req.LeftInterval),
		zap.String("right_interval:", req.RightInterval))

	// check interval valid
	sections, err := msection.Get(
		"group_id_right_interval >= %d AND group_id_left_interval <= %d",
		"", 0, 0, req.LeftInterval, req.RightInterval)
	if err != nil {
		zapwrapper.Logger.Error(
			"false to get section",
			zap.String("error", err.Error()),
		)
	}

	if len(sections) > 0 {
		// interval in request has been coverted
		// TODO(hackerzgz): return invalid interval error message
		return
	}

}
