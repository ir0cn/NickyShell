package service

import (
	"NickyShell/nicky/users/organization"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (svc *Service) ListOrganizations(ctx *gin.Context) {
	orgs, err := organization.ListOrganizations(svc.dbEngine)
	CheckError(err, fmt.Sprintf("ListOrganizations: %v", err))
	ctx.JSON(200, orgs)
}

func (svc *Service) NewOrganization(ctx *gin.Context) {

}
