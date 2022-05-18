package service

import (
	"NickyShell/nicky/users/organization"
	"github.com/gin-gonic/gin"
)

func (svc *Service) ListOrganizations(ctx *gin.Context) {
	orgs := organization.ListOrganizations(svc.dbEngine)
	ctx.JSON(200, orgs)
}

func (svc *Service) UpdateOrganization(ctx *gin.Context) {
	org := &organization.Organization{}
	ParamStruct(ctx, org)
	organization.UpdateOrganization(svc.dbEngine, org)
}

func (svc *Service) DeleteOrganization(ctx *gin.Context) {
	org := &organization.Organization{}
	ParamStruct(ctx, org)
	organization.DeleteOrganization(svc.dbEngine, org)
}
