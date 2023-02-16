package company_application_port_in

type IGetSiteInfoUserCase interface {
	GetSiteInfo(siteId int) string
}
