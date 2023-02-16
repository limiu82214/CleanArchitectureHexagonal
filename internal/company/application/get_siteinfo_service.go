package company_application

import (
	company_application_port_in "github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/in"
	company_domain "github.com/limiu82214/CleanArchitectureHexagonal/internal/company/domain"
)

type getSiteInfo struct {
}

func NewGetSiteInfo() company_application_port_in.IGetSiteInfoUserCase {
	return &getSiteInfo{}
}

func (gst *getSiteInfo) GetSiteInfo(siteID int) string {
	switch siteID {
	case 1:
		s := &company_domain.Site{Id: 1, SiteName: "site_name"}
		return s.SiteName
	case 2:
		s := &company_domain.Site{Id: 1, SiteName: "site_name2"}
		return s.SiteName
	case 3:
		s := &company_domain.Site{Id: 1, SiteName: "site_name2"}
		return s.SiteName
	}
	return "error"
}
