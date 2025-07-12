package controller

import (
	"Siu_shop/Backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	var product []model.Product
	var total int64
	//如果查询参数中没有 page，则使用默认值1
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	//如果查询参数中没有 pageSize，则使用默认值12
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "12"))
	//查询总数
	model.DB.Model(&model.Product{}).Count(&total)
	//分页查询
	if err := model.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&product).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"total":     total,
		"page":      page,
		"pageSize":  pageSize,
		"totalPage": int((total + int64(pageSize) - 1) / int64(pageSize)),
		"product":   product,
	})
}
