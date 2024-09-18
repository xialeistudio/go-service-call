package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	// 普通路由，发送STRING
	// 访问 http://localhost:8080/
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	// 普通路由，读取Query参数，发送JSON
	// 访问 http://localhost:8080/show?id=123
	r.GET("/show", func(c *gin.Context) {
		id := c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})
	// 普通路由，读取Path参数，发送JSON
	// 访问 http://localhost:8080/show/123
	r.GET("/show/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})
	userRoute := r.Group("/user")
	// 分组路由，读取表单参数，发送JSON
	// POST访问 http://localhost:8080/user/login
	userRoute.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	// 分组路由，读取JSON参数，发送JSON
	// POST访问 http://localhost:8080/user/register
	userRoute.POST("/register", func(c *gin.Context) {
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"password": user.Password,
			})
		}
	})
	// 中间件示例
	userRoute.GET("/profile", loginRequired(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "profile"})
	})
	log.Fatal(r.Run(":8080"))
}

func loginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization != "ok" {
			// 阻断请求
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		// 继续执行后续Handler
		c.Next()
	}
}
