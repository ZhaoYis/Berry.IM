package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (bc *BaseController) Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"data": nil,
		"msg":  msg})
}

func (bc *BaseController) Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}

func (bc *BaseController) SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}
