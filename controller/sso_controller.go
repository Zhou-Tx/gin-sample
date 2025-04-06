package controller

import (
	"gin-sample/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SsoController struct{}

func (c SsoController) RegisterRoutes(engine *gin.Engine) {
	group := engine.Group("/sso")
	group.POST("/session", loginHandler)
	//userGroup.GET("/status", StatusHandler)
}

func loginHandler(ctx *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 调用 Service 层进行登录验证
	err := service.Login(loginData.Username, loginData.Password, ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

//
//func StatusHandler(c *gin.Context) {
//	username, _ := c.Get("username")
//	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "username": username})
//}
