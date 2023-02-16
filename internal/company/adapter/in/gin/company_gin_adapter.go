package company_gin_adapter_in

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type companyGinAdapter struct {
}

func NewCompanyGinAdapter() *companyGinAdapter {
	return &companyGinAdapter{}
}

func (cga *companyGinAdapter) GetSiteInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "getSiteInfo",
	})
}
