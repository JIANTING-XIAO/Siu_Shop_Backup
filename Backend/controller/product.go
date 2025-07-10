package controller

import (
	"Siu_shop/Backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	var product []model.Product
	if err := model.DB.Find(&product).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"total":   len(product),
		"product": product,
	})
}
