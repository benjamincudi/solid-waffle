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

type productCreateControllerDS interface {
	GetVariations(ctx context.Context) ([]models.Variation, error)
	InsertVariation(ctx context.Context, name string) (models.Variation, error)
	InsertProductVariation(ctx context.Context, variation models.ProductVariation) (models.ProductVariation, error)
	InsertProduct(ctx context.Context, sellerId int, name string) (models.Product, error)
}

type ProductCreateController struct {
	DS productCreateControllerDS
}

type variationForm struct {
	Variations  []string `json:"variations"`
	Price       int      `json:"price"`
	IsAvailable bool     `json:"is_available"`
}

type productCreateForm struct {
	Name       string          `json:"name"`
	SellerID   int             `json:"seller_id"`
	Variations []variationForm `json:"variations"`
}

func (ctrl ProductCreateController) HandlePost(c *gin.Context) {
	ctx := c.Request.Context()
	var form productCreateForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to bind product form"})
		return
	}
	variationToID := map[string]int{}
	for _, v := range form.Variations {
		for _, variation := range v.Variations {
			variationToID[variation] = -1
		}
	}
	existingVariations, err := ctrl.DS.GetVariations(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, existingVariation := range existingVariations {
		if _, ok := variationToID[existingVariation.Name]; ok {
			variationToID[existingVariation.Name] = existingVariation.ID
		}
	}
	for name, id := range variationToID {
		if id == -1 {
			v, err := ctrl.DS.InsertVariation(ctx, name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			variationToID[name] = v.ID
		}
	}
	product, err := ctrl.DS.InsertProduct(ctx, form.SellerID, form.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var newVariations []models.ProductVariation
	for _, v := range form.Variations {
		var variationIDs []int
		for _, name := range v.Variations {
			variationIDs = append(variationIDs, variationToID[name])
		}
		variation, err := ctrl.DS.InsertProductVariation(ctx, models.ProductVariation{
			ProductID:    product.ID,
			VariationIDs: variationIDs,
			IsAvailable:  v.IsAvailable,
			Price:        v.Price,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		newVariations = append(newVariations, variation)
	}
	c.JSON(http.StatusOK, newVariations)
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
