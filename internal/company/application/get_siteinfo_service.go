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

func (gsi *getSiteInfo) GetSiteInfo(siteID int) string {
	s := gsi.loadCompanyPort.GetSite(siteID)
	switch siteID {
	case 1:
		return s.SiteName
	case 2:
		return s.SiteName
	case 3:
		return s.SiteName
	}
	return "error"
}
