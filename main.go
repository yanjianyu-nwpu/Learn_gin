package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	limitPtr *rate.Limiter
)

func rateLimitMiddleWare(c *gin.Context) {
	//if limitPtr.Allow() {

	//	c.Next()
	//	return
	//}
	//c.String(http.StatusOK, "rate limit ...")
	//c.Abort()

	limitPtr.Wait(c.Request.Context())
	c.Next()
}
func main() {
	limitPtr = rate.NewLimiter(1, 10)
	engine := gin.Default()
	engine.Use(rateLimitMiddleWare)
	//l := rate.NewLimiter(100.0, 100)
	//fmt.Println(l)
}
