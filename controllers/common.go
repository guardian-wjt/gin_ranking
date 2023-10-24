package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

// 统一Json响应格式
type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
type JsonErrorStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// 正确
func ReturnSuccess(ctx *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	ctx.JSON(200, json)
}

// 错误
func ReturnError(ctx *gin.Context, code int, msg interface{}) {
	json := &JsonErrorStruct{Code: code, Msg: msg}
	ctx.JSON(200, json)
}

// md5加密
func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
