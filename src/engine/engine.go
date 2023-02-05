package engine

import (
	"GoFun/src/controller"
	"GoFun/src/middleware"
	"github.com/gin-gonic/gin"
)

func Init() (engine *gin.Engine) {
	engine = gin.Default()

	engine.LoadHTMLGlob("public/templates/*")
	engine.Static("/static", "public/static")

	// 首页
	engine.GET("/", middleware.Set(), controller.Index)
	// 获取新的毒鸡汤句子
	engine.GET("/getFunWords", middleware.Set(), controller.GetFunWords)

	return engine
}
