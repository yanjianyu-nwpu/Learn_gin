package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yanjianyu-nwpu/Learn_gin/controllers"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	//注册：
	router.GET("/register", controllers.RegisterGet)
	return router

}
