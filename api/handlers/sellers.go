package handlers

import (
	"context"
	"net/http"
	"solid-waffle/models"

	"github.com/gin-gonic/gin"
)

type sellersControllerDS interface {
	GetSellers(ctx context.Context) ([]models.Seller, error)
	InsertSeller(ctx context.Context, name string) (models.Seller, error)
}

type SellersController struct {
	DS sellersControllerDS
}

func (ctrl SellersController) HandleGet(c *gin.Context) {
	sellers, err := ctrl.DS.GetSellers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sellers)
}

type sellerForm struct {
	Name string `form:"name" json:"name"`
}

func (ctrl SellersController) HandlePost(c *gin.Context) {
	var newSeller sellerForm
	if err := c.ShouldBind(&newSeller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to bind seller form"})
		return
	}
	if len(newSeller.Name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name must be non-empty string"})
		return
	}
	seller, err := ctrl.DS.InsertSeller(c.Request.Context(), newSeller.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seller)
}
