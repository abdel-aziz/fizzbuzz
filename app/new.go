package app

import (
	"fmt"
)

func NewApp(config Config) (App, error) {
	var app App
	switch config.Router {
	case ROUTER_GOCRAFT:
		app = NewAppGoCraft()
	// case ROUTER_DEFAULT_TESTER:
	// 	app = NewDefaultAppTester()
	default:
		return nil, fmt.Errorf("Router not implemented: %s", config.Router)
	}

	app.InitMiddlewares(config)
	app.InitContext(config)

	return app, nil
}
