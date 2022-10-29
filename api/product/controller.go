package product

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/exceptions"
)

type controllerProduct struct {
	service productService
}

func ControllerProduct() controllerProduct {
	return controllerProduct{
		service: ServiceProduct(),
	}
}

func (p controllerProduct) SaveProduct(c *gin.Context) {
	user := getUser(c)

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
	user := getUser(c)
	allProduct, err := p.service.FindAllProduct(c, user)
	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
	c.JSON(http.StatusOK, allProduct)
	return
}

func (p controllerProduct) DeleteProduct(c *gin.Context) {
	user := getUser(c)
	productId := c.Param("productId")
	if err := p.service.DeleteProductById(c, productId, user); err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}
}

func getUser(c *gin.Context) domain.LoggedUser {
	loggedUser, _ := c.Get("loggedUser")
	return loggedUser.(domain.LoggedUser)
}
