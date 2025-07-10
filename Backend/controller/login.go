package controller

import (
	"Siu_shop/Backend/model"
	"Siu_shop/Backend/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var login LoginRequest
	var user model.User
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求数据格式错误",
		})
		return
	}
	//验证账号密码是否为空
	if login.Account == "" || login.Password == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":     http.StatusUnprocessableEntity,
			"password": login.Password,
			"account":  login.Account,
			"message":  "用户名或密码不能为空",
		})
		return
	}

	//验证账号是否存在
	result := model.DB.Where("account=?", login.Account).Find(&user)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户不存在",
		})
		return
	}

	//验证账号密码是否正确
	if login.Password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "密码错误",
		})
		return
	}

	//生成token给他
	token, err := tool.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "生成token失败",
		})
		return
	}

	//登录成功
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
		"token":   gin.H{"token": token},
	})
}
