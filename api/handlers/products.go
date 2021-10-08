package handlers

import (
	"context"
	"net/http"
	"solid-waffle/models"

	"github.com/gin-gonic/gin"
)

type productsControllerDS interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	InsertProduct(ctx context.Context, sellerId int, name string) (models.Product, error)
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

func (ctrl ProductsController) HandlePost(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to bind seller form"})
		return
	}
	product, err := ctrl.DS.InsertProduct(c.Request.Context(), product.SellerID, product.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
type productVariationsControllerDS interface {
	GetProductVariations(ctx context.Context) ([]models.ProductVariation, error)
	InsertProductVariation(ctx context.Context, variation models.ProductVariation) (models.ProductVariation, error)
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

func (ctrl ProductVariationsController) HandlePost(c *gin.Context) {
	var v models.ProductVariation
	if err := c.Bind(&v); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	variation, err := ctrl.DS.InsertProductVariation(c.Request.Context(), v)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, variation)
}
