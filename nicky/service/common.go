package service

import (
	. "NickyShell/nicky/exception"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

/****************************************************
 * 请求处理
 ****************************************************/

// ParamStruct 从request.Body读，并组织成struct
func ParamStruct(ctx *gin.Context, obj interface{}) {
	defer func() { _ = ctx.Request.Body.Close() }()
	data, err := ioutil.ReadAll(ctx.Request.Body)
	ThrowIfError(err)
	err = json.Unmarshal(data, obj)
	ThrowIfError(err)
}
