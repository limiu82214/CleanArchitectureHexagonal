package company_application_port_out

import (
	company_domain "github.com/limiu82214/CleanArchitectureHexagonal/internal/company/domain"
)

type ILoadCompanyPort interface {
	GetSite(siteID int) company_domain.Site
}
