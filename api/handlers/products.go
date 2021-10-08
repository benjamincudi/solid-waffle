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
