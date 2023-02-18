package in

import (
	"errors"
)

type GetSiteInfoCommand struct {
	SiteID int
}

var errSiteIDNegative = errors.New("siteID must bigger than 0")

func NewGetSiteInfoCommand(siteID int) (*GetSiteInfoCommand, error) {
	if siteID < 1 {
		return nil, errSiteIDNegative
	}

	return &GetSiteInfoCommand{
		SiteID: siteID,
	}, nil
}
