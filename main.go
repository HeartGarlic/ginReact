package main

import (
	"errors"
	"ginReact/conf"
	"ginReact/router"
	"github.com/gin-gonic/gin"
)

// main 程序入口
func main() {
	// 初始化gin
	r := gin.Default()
	// 初始化配置文件 / 数据库连接
	conf.InitConf()
	// 初始化路由
	router.InitRouter(r)
	// 加载模板
	r.LoadHTMLGlob("views/*/*")
	// 静态资源
	r.Static("/public", "./public")
	// 启动程序
	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
