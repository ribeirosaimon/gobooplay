package subscription

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ribeirosaimon/gobooplay/exceptions"
	"ribeirosaimon/gobooplay/util"
)

type controllerSubsription struct {
	service SubscriptionService
}

func ControllerProduct() controllerSubsription {
	return controllerSubsription{
		service: ServiceSubscription(),
	}
}

func (s controllerSubsription) findMySubscription(c *gin.Context) {
	user := util.GetUser(c)
	subscription, err := s.service.findSubscription(c, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusOK, subscription)
	return
}

func (s controllerSubsription) validationSubscription(c *gin.Context) {
	user := util.GetUser(c)
	if err := s.service.ValidateSubscription(c, user); err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}

	c.JSON(http.StatusOK, "ok")
	return
}
