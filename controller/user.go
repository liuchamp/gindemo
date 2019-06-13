package controller

import (
	"github.com/detectiveHLH/go-backend-starter/util"
	"github.com/gin-gonic/gin"
	"p1/middleware/consts"
)

// @Summary 用户详情
// @Produce  json
// @Param username query string true "username"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/user [get]
func Info(c *gin.Context) {
	appG := util.Gin{C: c}
	username := c.Query("username")
	if username == "" || len(username) == 0 {
		appG.Response(200, consts.INVALIDPARAMS, "username is empty")
	}
	c.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"message": "This works",
		"data":    nil,
	})
}
