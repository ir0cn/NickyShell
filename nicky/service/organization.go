package service

import (
	"NickyShell/nicky/users/organization"
	"github.com/gin-gonic/gin"
)

func (svc *Service) ListOrganizations(ctx *gin.Context) {
	orgs, err := organization.ListOrganizations(svc.dbEngine)
	CheckError(err, &Response{1001, "", nil})
	ctx.JSON(200, &Response{Data: orgs})
}

func (svc *Service) NewOrganization(ctx *gin.Context) {

}
