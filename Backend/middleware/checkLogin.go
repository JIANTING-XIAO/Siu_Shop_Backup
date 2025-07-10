package middleware

import "github.com/gin-gonic/gin"

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, err := c.Cookie("name")
		if err != nil || name == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "用户未登录",
			})
			c.Abort()
		}
		c.Set("userName", name)
		c.Next()
	}
}
