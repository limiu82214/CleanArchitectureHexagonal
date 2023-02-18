package in

import "fmt"

type GetSiteInfoCommand struct {
	SiteID int
}

func NewGetSiteInfoCommand(siteID int) (*GetSiteInfoCommand, error) {
	if siteID < 1 {
		return nil, fmt.Errorf("siteID must bigger than 0")
	}
	return &GetSiteInfoCommand{
		SiteID: siteID,
	}, nil
}
