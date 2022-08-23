package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	BaseController
}

// Index 前端页面渲染入口
func (i IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index/index.html", gin.H{})
}
