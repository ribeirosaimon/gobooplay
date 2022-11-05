package product

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/exceptions"
	"ribeirosaimon/gobooplay/util"
)

type controllerProduct struct {
	service ProductService
}

func ControllerProduct() controllerProduct {
	return controllerProduct{
		service: ServiceProduct(),
	}
}

func (p controllerProduct) saveProduct(c *gin.Context) {
	user := util.GetUser(c)

	var payload domain.ProductDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		exceptions.ValidateException(c, "incorrect body", http.StatusConflict)
		return
	}
	product, err := p.service.AddProduct(c, payload, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, product)
	return
}

func (p controllerProduct) FindAvailableProduct(c *gin.Context) {
	user := util.GetUser(c)
	allProduct, err := p.service.FindAllProduct(c, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusOK, allProduct)
	return
}

func (p controllerProduct) deleteProduct(c *gin.Context) {
	user := util.GetUser(c)
	productId := c.Param("productId")
	if err := p.service.DeleteProductById(c, productId, user); err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	return
}

func (p controllerProduct) updateProduct(c *gin.Context) {
	user := util.GetUser(c)
	productId := c.Param("productId")

	var payload domain.ProductDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		exceptions.ValidateException(c, "incorrect body", http.StatusConflict)
		return
	}
	product, err := p.service.UpdateProduct(c, payload, productId, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusCreated, product)
	return
}
