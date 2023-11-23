package handler

import (
	"net/http"
	"strconv"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/service"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{productService: productService}
}

func (ph *productHandler) CreateProduct(c *gin.Context) {
	var p dto.NewProductRequest

	if err := c.ShouldBindJSON(&p); err != nil {
		errBind := errs.NewUnprocessibleEntityError("invalid json request body")
		c.AbortWithStatusJSON(errBind.Status(), errBind)
		return
	}

	resp, err := ph.productService.CreateProduct(p)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (ph *productHandler) GetProduct(c *gin.Context) {
	resp, err := ph.productService.GetProducts()
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, resp.Data)
}

func (ph *productHandler) UpdateProduct(c *gin.Context) {
	var p dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		errBind := errs.NewUnprocessibleEntityError("invalid json request body")
		c.AbortWithStatusJSON(errBind.Status(), errBind)
		return
	}

	param := c.Param("productId")
	productId, _ := strconv.Atoi(param)
	p.ID = productId

	resp, err := ph.productService.UpdateProduct(p)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (ph *productHandler) DeleteProduct(c *gin.Context) {
	param := c.Param("productId")
	productId, _ := strconv.Atoi(param)

	resp, err := ph.productService.DeleteProduct(productId)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
