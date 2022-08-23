package api

import (
	"ginReact/controllers"
	"ginReact/models"
	"ginReact/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
    controllers.BaseController
}

//AddUser 新增用户接口
func (u UserController) AddUser(c *gin.Context)  {
    user   := models.NewUser()
	if err := c.ShouldBind(user); err != nil {
        c.JSON(http.StatusOK, u.Error(err.Error()))
        return
    }
	// 获取用户登陆Ip
	user.LastLoginIp = (&tool.Tool{}).IpToInt(c.ClientIP())
	add, err := user.Add()
	if err != nil {
		c.JSON(http.StatusOK, u.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, u.Success(map[string]int64{"id": add}))
	return
}

// Login 登陆接口
func (u UserController) Login(c *gin.Context){
	loginUser := &models.LoginUser{}
	err := c.ShouldBind(loginUser)
	if err != nil {
		c.JSON(http.StatusOK, u.Error(err.Error()))
		return
	}
	user := models.NewUser()
	// 将校验过得值赋值给user对象
	user.Username = loginUser.Username
	user.Email    = loginUser.Email
	user.Password = loginUser.Password
	user.Phone    = loginUser.Phone
	login, err := models.NewUser().Login(user)
	if err != nil {
		c.JSON(http.StatusOK, u.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, u.Success(map[string]*models.User{"userInfo": login}))
	return
}
