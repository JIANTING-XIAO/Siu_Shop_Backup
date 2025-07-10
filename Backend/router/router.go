package router

import (
	"Siu_shop/Backend/controller"
	"Siu_shop/Backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func InitGin() {
	r := gin.Default()
	//cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//路由
	r.GET("/userinfo", middleware.AuthMiddleware(), controller.GetUserInfo)
	r.GET("/product", controller.GetAllProduct)
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)

	//启动端口
	if err := r.Run(":8080"); err != nil {
		panic("启动失败, error=" + err.Error())
	}
}
