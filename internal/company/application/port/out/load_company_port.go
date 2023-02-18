package out

import (
	"github.com/limiu82214/CleanArchitectureHexagonal/internal/company/domain"
)

type ILoadCompanyPort interface {
	GetSite(siteID int) domain.Site
}
