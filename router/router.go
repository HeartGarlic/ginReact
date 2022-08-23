package router

import (
	"ginReact/controllers"
	"ginReact/controllers/api"
	"ginReact/middlewares"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	// 跨域中间件
	r.Use(middlewares.Cors())
	// 首页方法
	r.GET("/", controllers.IndexController{}.Index)

	// 接口方法
	ApiGroup := r.Group("/api")
	{
		// v1 版本接口
		v1 := ApiGroup.Group("/v1")
		{
			// 用户组
			user := v1.Group("/user")
			{
				// 注册
				user.POST("/add", api.UserController{}.AddUser)
				// 登陆
				user.POST("/login", api.UserController{}.Login)
			}
			// 帖子
			blog := v1.Group("/blog")
			{
				// 发布
				blog.POST("/add", api.BlogController{}.AddBlog)
				// 修改
				// 删除
				// 查看
				// 点赞

			}
			// 评论
			// comment := v1.Group("/comment")
			{
				// 发布
				// 修改
				// 删除
				// 查看
				// 点赞
				// 回复
			}
		}
	}
}
