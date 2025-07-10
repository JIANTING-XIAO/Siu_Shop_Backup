package controller

import (
	"Siu_shop/Backend/dto"
	"Siu_shop/Backend/model"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"msg":  "获取用户信息成功",
		"user": dto.ToUserDto(user.(model.User)),
	})
}
