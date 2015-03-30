package config

import (
	"github.com/naoina/kocha"
	"github.com/naoina/vocanew/app/controller"
)

type RouteTable kocha.RouteTable

var routes = RouteTable{
	{
		Name:       "root",
		Path:       "/",
		Controller: &controller.Root{},
	}, {
		Name:       "about",
		Path:       "/about",
		Controller: &controller.About{},
	}, {
		Name:       "contact",
		Path:       "/contact",
		Controller: &controller.Contact{},
	}, {
		Name:       "newsong",
		Path:       "/newsong",
		Controller: &controller.Newsong{},
	},
}

func init() {
	AppConfig.RouteTable = kocha.RouteTable(append(routes, RouteTable{
		{
			Name:       "static",
			Path:       "/*path",
			Controller: &kocha.StaticServe{},
		},
	}...))
}
