package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouterFunc func(string, ...gin.HandlerFunc) gin.IRoutes
type HandleFunc func(*gin.Context)

type Router struct {
	Uri      string
	Method   RouterFunc
	Action   HandleFunc
	NeedAuth bool
}

// makeHandler
func (svc *Service) makeHandler(f HandleFunc, needAuth bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if res, ok := r.(*Response); ok {
					ctx.JSON(400, res)
				} else {
					ctx.JSON(400, &Response{Code: 1000, Message: fmt.Sprintf("%v", r)})
				}
				ctx.Abort()
				return
			}
		}()

		if needAuth && !svc.Verify(ctx) { // 检查认证信息
			return
		}

		f(ctx)
	}
}

// StartWebService 启动Web服务
func (svc *Service) StartWebService() error {
	router := gin.Default()
	var routers = []Router{
		{"/api/v1/user/:userName/login", router.POST, svc.UserLogin, false},
		{"/api/v1/assets/:assetId/connect", router.GET, svc.WsConnect, false},

		// 组织
		{"/api/v1/organization", router.GET, svc.ListOrganizations, false},
		{"/api/v1/organization", router.POST, svc.NewOrganization, false},
	}
	for _, r := range routers {
		r.Method(r.Uri, svc.makeHandler(r.Action, r.NeedAuth))
	}
	return router.Run(":8081")
}
