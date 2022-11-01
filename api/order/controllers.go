package order

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/util"
)

type controllerShoop struct {
	service shoopService
}

func ControllerShoop() controllerShoop {
	return controllerShoop{
		service: ServiceShoop(),
	}
}
func (s controllerShoop) SendOrder(c *gin.Context) {
	user := util.GetUser(c)
	s.service.sendOrder(c, user)
}
