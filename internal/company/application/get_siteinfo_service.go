package company_application

import (
	company_application_port_in "github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/in"
	company_application_port_out "github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/out"
)

type getSiteInfo struct {
	loadCompanyPort company_application_port_out.ILoadCompanyPort
}

func NewGetSiteInfo(loadCompanyPort company_application_port_out.ILoadCompanyPort) company_application_port_in.IGetSiteInfoUserCase {
	return &getSiteInfo{
		loadCompanyPort: loadCompanyPort,
	}
}

func (gsi *getSiteInfo) GetSiteInfo(gsic company_application_port_in.GetSiteInfoCommand) string {
	s := gsi.loadCompanyPort.GetSite(gsic.SiteID)
	return s.SiteName
}
