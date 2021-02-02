package controller

import "github.com/gin-gonic/gin"

type UserController struct {}

func UserRegister(group *gin.RouterGroup) {
	user := &UserController{}
	group.POST("/reg", user.reg)
	group.POST("/sendcode", user.sendCode)
	group.POST("/phonelogin", user.phoneLogin)
	group.POST("/otherlogin", user.otherLogin)
}

func UserCheckRegister(group *gin.RouterGroup) {
	user := &UserController{}
	group.POST("/logout", user.logout)
}

// reg godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户接口
// @ID /user/reg
// @Accept  json
// @Produce  json
// @Param body body dto.UserInput true "body"
// @Success 200 {object} middleware.Response{data=dto.UserOutput} "success"
// @Router /user/reg [post]
func (user *UserController) reg(c *gin.Context) {

}

func (user *UserController) sendCode(c *gin.Context) {

}

func (user *UserController) phoneLogin(c *gin.Context) {
	
}

func (user *UserController) otherLogin(c *gin.Context) {

}

func (user *UserController) logout(c *gin.Context) {

}