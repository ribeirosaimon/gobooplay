package voucher

import (
	"github.com/gin-gonic/gin"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/middleware"
)

func RouteVoucher(e *gin.RouterGroup) {
	group := e.Group("/voucher")

	group.Use(middleware.Authorization([]domain.Role{domain.ADMIN})).POST("/", ControllerVoucher().createVoucher)
	group.Use(middleware.Authorization([]domain.Role{domain.ADMIN})).PUT("/:voucherId", ControllerVoucher().updateVoucher)
	group.Use(middleware.Authorization([]domain.Role{domain.ADMIN})).DELETE("/:voucherId", ControllerVoucher().deleteVoucher)
}
