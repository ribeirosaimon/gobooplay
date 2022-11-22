package voucher

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/exceptions"
	"ribeirosaimon/gobooplay/util"
)

type controllerVoucher struct {
	service voucherService
}

func ControllerVoucher() controllerVoucher {
	return controllerVoucher{
		service: ServiceVoucher(),
	}
}

func (s controllerVoucher) createVoucher(c *gin.Context) {
	user := util.GetUser(c)

	var payload domain.VoucherDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		exceptions.ValidateException(c, "incorrect body", http.StatusConflict)
		return
	}
	product, err := s.service.AddVoucher(c, payload, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, product)
	return
}

func (s controllerVoucher) deleteVoucher(c *gin.Context) {
	user := util.GetUser(c)
	voucherId := c.Param("voucherId")
	if err := s.service.deleteVoucherById(c, voucherId, user); err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	return
}

func (s controllerVoucher) updateVoucher(c *gin.Context) {
	user := util.GetUser(c)
	voucherId := c.Param("voucherId")

	var payload domain.VoucherDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		exceptions.ValidateException(c, "incorrect body", http.StatusConflict)
		return
	}
	voucher, err := s.service.updateVoucher(c, payload, voucherId, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, voucher)
	return
}

func (s controllerVoucher) getVoucher(c *gin.Context) {
	voucherId := c.Param("voucherId")
	voucher, err := s.service.getVoucher(c, voucherId)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, voucher)
	return
}
