package main

import (
	"github.com/gin-gonic/gin"
	"my/blogs/controller"
	"my/blogs/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) //中间件用于保护用户信息接口,如果没有使用中间件，返回的user就是null,因为上下文没有存这个信息
	return r
}
