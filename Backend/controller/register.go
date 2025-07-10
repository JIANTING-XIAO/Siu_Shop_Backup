package controller

import (
	"Siu_shop/Backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"time"
)

type RegisterRequest struct {
	Account   string `json:"account" form:"account"`
	Password  string `json:"password" form:"password"`
	Telephone string `json:"telephone" form:"telephone"`
}

func Register(c *gin.Context) {
	var register RegisterRequest
	//先验证请求数据格式
	if err := c.ShouldBind(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求数据格式错误",
		})
		return
	}
	//验证账号密码手机号是否为空
	if register.Account == "" || register.Telephone == "" || register.Password == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "账号、密码或手机号不能为空",
		})
		return
	}
	//验证手机号是否是11位数
	if len(register.Telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "手机号必须是11位数",
		})
		return
	}
	//验证密码是否符合规则：密码只能包含字母和数字，且长度为6到15位
	passwordRule := `^[a-zA-Z\d]{6,15}$`
	matchPassword, _ := regexp.MatchString(passwordRule, register.Password)
	if !matchPassword {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "密码只能包含字母和数字，且长度为6到15位",
		})
		return
	}

	//验证账号是否符合规则
	accountRule := `^[a-zA-Z\d]{6,15}$`
	matchAccount, _ := regexp.MatchString(accountRule, register.Account)
	if !matchAccount {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "账号只能包含字母和数字，且长度为6到15位",
		})
		return
	}
	//验证手机号是否存在
	if telephoneExist(model.DB, register.Telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "手机号已存在",
		})
		return
	}
	//验证账号是否已存在
	if accountExist(model.DB, register.Account) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "账号已存在",
		})
		return
	}
	//所有验证通过，创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "密码加密失败",
		})
	}
	newUser := model.User{
		Account:     register.Account,
		Password:    string(hashPassword),
		Telephone:   register.Telephone,
		CreatedTime: time.Now(),
	}
	createResult := model.DB.Create(&newUser)
	if createResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to create user",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func telephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	result := db.Table("user").Where("telephone =?", telephone).First(&user)
	if result.RowsAffected > 0 {
		// 表示存在相同的 telephone
		fmt.Println("手机号已存在")
		return true
	} else {
		// 表示不存在相同的 telephone
		fmt.Println("手机号不存在,可以使用该手机号注册")
		return false
	}
}
func accountExist(db *gorm.DB, account string) bool {
	var user model.User
	result := db.Table("user").Where("account =?", account).First(&user)
	if result.RowsAffected > 0 {
		// 表示存在相同的 telephone
		fmt.Println("账号已存在")
		return true
	} else {
		// 表示不存在相同的 telephone
		fmt.Println("账号不存在,可以使用该手机号注册")
		return false
	}
}
