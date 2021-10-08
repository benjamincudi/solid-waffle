package handlers

import (
	"context"
	"net/http"
	"solid-waffle/models"

	"github.com/gin-gonic/gin"
)

type productsControllerDS interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
}

type ProductsController struct {
	DS productsControllerDS
}

func (ctrl ProductsController) HandleGet(c *gin.Context) {
	products, err := ctrl.DS.GetProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

type productVariationsControllerDS interface {
	GetProductVariations(ctx context.Context) ([]models.ProductVariation, error)
}

type ProductVariationsController struct {
	DS productVariationsControllerDS
}

func (ctrl ProductVariationsController) HandleGet(c *gin.Context) {
	variations, err := ctrl.DS.GetProductVariations(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, variations)
}
