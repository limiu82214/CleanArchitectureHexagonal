package company_gin_adapter_in

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	company_application_port_in "github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/in"
)

type companyGinAdapter struct {
	getSiteInfoUseCase company_application_port_in.IGetSiteInfoUserCase
}

func NewCompanyGinAdapter(getSiteInfoUseCase company_application_port_in.IGetSiteInfoUserCase) *companyGinAdapter {
	return &companyGinAdapter{
		getSiteInfoUseCase: getSiteInfoUseCase,
	}
}

func (cga *companyGinAdapter) GetSiteInfo(ctx *gin.Context) {
	siteId, err := strconv.Atoi(ctx.DefaultQuery("site_id", "1"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": cga.getSiteInfoUseCase.GetSiteInfo(siteId),
	})
}
