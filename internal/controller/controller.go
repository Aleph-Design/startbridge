package controller

import (
	"github.com/aleph-design/startbridge/internal/bidmap"
	"github.com/aleph-design/startbridge/internal/config"
	"github.com/aleph-design/startbridge/internal/models"
)

// site wide access to app config
var app *config.AppConfig

// actually populate this package with site-wide
// access to AppConfig by transferring a pointer
// *config.AppConfig and call this function from
// main()
func NewController(a *config.AppConfig) {
	app = a
}

func CheckValidBid(plr int, bid *bidmap.Bid, g *models.Game) bool {

	return false
}
