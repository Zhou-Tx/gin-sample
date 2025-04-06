package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	RegisterRoutes()
}

func RegisterRoutes(engine *gin.Engine) {
	SsoController{}.RegisterRoutes(engine)
}
