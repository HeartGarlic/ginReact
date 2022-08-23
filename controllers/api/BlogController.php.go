package api

import (
	"ginReact/controllers"
	"ginReact/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogController struct {
	controllers.BaseController
}

// AddBlog 新增帖子
func (b BlogController) AddBlog(c *gin.Context) {
	blog := models.NewBlog()
	err := c.ShouldBind(blog)
	if err != nil {
		c.JSON(http.StatusOK, b.Error(err.Error()))
		return
	}
	// 校验用户是否存在
	user := models.NewUser().FindUserOfId(blog.Uid)
	if user == nil {
		c.JSON(http.StatusOK, b.Error("用户不存在"))
        return
	}
	addBlog, err := blog.AddBlog()
	if err != nil || !addBlog {
		c.JSON(http.StatusOK, b.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, b.Success(nil))
	return
}

// List 帖子列表
func (b BlogController) List(c *gin.Context){
	// page := c.Query("page")
	// limit:= c.Query("limit")
}

// Detail 帖子详情
func(b BlogController) Detail(c *gin.Context){
	// id := c.Query("id")
}

