package gin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/in"
)

type companyGinAdapter struct {
	getSiteInfoUseCase in.IGetSiteInfoUserCase
}

func NewCompanyGinAdapter(getSiteInfoUseCase in.IGetSiteInfoUserCase) *companyGinAdapter {
	return &companyGinAdapter{
		getSiteInfoUseCase: getSiteInfoUseCase,
	}
}

func (cga *companyGinAdapter) GetSiteInfo(ctx *gin.Context) {
	siteID, err := strconv.Atoi(ctx.DefaultQuery("site_id", "1"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	getSiteInfoCommand := in.GetSiteInfoCommand{
		SiteID: siteID,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": cga.getSiteInfoUseCase.GetSiteInfo(getSiteInfoCommand),
	})
}
