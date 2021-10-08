package handlers

import (
	"context"
	"net/http"
	"solid-waffle/models"

	"github.com/gin-gonic/gin"
)

type sellersControllerDS interface {
	GetSellers(ctx context.Context) ([]models.Seller, error)
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
