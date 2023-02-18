package application

import (
	"github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/in"
	"github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/out"
)

type getSiteInfo struct {
	loadCompanyPort out.ILoadCompanyPort
}

func NewGetSiteInfo(loadCompanyPort out.ILoadCompanyPort) in.IGetSiteInfoUserCase {
	return &getSiteInfo{
		loadCompanyPort: loadCompanyPort,
	}
}

func (gsi *getSiteInfo) GetSiteInfo(gsic in.GetSiteInfoCommand) string {
	s := gsi.loadCompanyPort.GetSite(gsic.SiteID)
	return s.SiteName
}
