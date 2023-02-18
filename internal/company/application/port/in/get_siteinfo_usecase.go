package in

type IGetSiteInfoUserCase interface {
	GetSiteInfo(gsic GetSiteInfoCommand) string
}
