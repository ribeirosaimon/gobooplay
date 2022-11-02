package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ribeirosaimon/gobooplay/exceptions"
	"ribeirosaimon/gobooplay/util"
)

type controllerShoop struct {
	service OrderService
}

func ControllerShoop() controllerShoop {
	return controllerShoop{
		service: ServiceOrder(),
	}
}
func (s controllerShoop) SendOrder(c *gin.Context) {
	user := util.GetUser(c)
	order, err := s.service.sendOrder(c, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}

	c.JSON(http.StatusOK, order)
	return
}
