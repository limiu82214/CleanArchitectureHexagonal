package company_application

import (
	company_application_port_in "github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/in"
)

type getSiteInfo struct {
}

func NewGetSiteInfo() company_application_port_in.IGetSiteInfoUserCase {
	return &getSiteInfo{}
}

func (gst *getSiteInfo) GetSiteInfo(siteID int) string {
	switch siteID {
	case 1:
		return "site_name"
	case 2:
		return "site_name2"
	case 3:
		return "site_name3"
	}
	return "error"
}
