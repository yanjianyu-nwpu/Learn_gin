package main

import "github.com/yanjianyu-nwpu/Learn_gin/routers"

func main() {

	router := routers.InitRouter()
	// 静态资源
	router.Static("/static", "./static")
	router.Run(":8080")
}
