package dto

import (
	"github.com/e421083458/go_chat/public"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	UserName string `json:"username" form:"username" comment:"用户名" example:"admin" validate:"required"`// 用户名
}

type UserOutput struct {
	Token	string	`json:"token" form:"token" comment:"token" example:"token" validate:""`//token
}

func (param *UserInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}