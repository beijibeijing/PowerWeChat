package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

func Inject(app kernel.ApplicationInterface) {
	container := app.GetContainer()
	container.GetConfig()

	component := &Client{
		kernel.NewBaseClient(&app, nil),
	}
	components := app.GetComponents()

	(*components)["base"] = component
}