package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGet(c *gin.Context) {
	//c.String(http.StatusOK, "is ok ")

	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}
