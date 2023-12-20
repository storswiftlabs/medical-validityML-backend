package handlers

import (
	"github.com/gin-gonic/gin"
)

type BaseResp struct {
	Ok      bool   `json:"ok"`
	Count   int64  `json:"count"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

type BaseRespNoCount struct {
	Ok      bool   `json:"ok"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func returnData(code int, count int64, msg string, data any, c *gin.Context) {
	c.JSON(code, BaseResp{
		Ok:      true,
		Count:   count,
		Message: msg,
		Data:    data,
	})
}

func returnDataNoCount(code int, msg string, data any, c *gin.Context) {
	c.JSON(code, BaseRespNoCount{
		Ok:      true,
		Message: msg,
		Data:    data,
	})
}

type Progress struct {
	Progress int    `json:"progress"`
	Message  string `json:"message"`
	Error    error  `json:"error"`
}

func returnError(code int, message string, err error, c *gin.Context) {
	c.JSON(code, BaseResp{
		Ok:      false,
		Count:   0,
		Message: message,
	})
	if err != nil {
		_ = c.Error(err)
	}
}
